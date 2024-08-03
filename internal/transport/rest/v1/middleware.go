package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodeyeen/chatsy/api/v1"
)

func NewCheckJWTMiddleware(secret string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("accessToken")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(api.ErrorResponse{
					Message: "accessToken was not provided",
				})
				return
			}

			token, err := jwt.ParseWithClaims(cookie.Value, &api.TokenClaims{}, func(token *jwt.Token) (any, error) {
				return []byte(secret), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(api.ErrorResponse{
					Message: "malformed accessToken",
				})
				return
			}

			claims, ok := token.Claims.(*api.TokenClaims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(api.ErrorResponse{
					Message: "unknown claims type, cannot proceed",
				})
				return
			}

			ctx := context.WithValue(r.Context(), "userID", claims.UserID)

			next(w, r.WithContext(ctx))
		}
	}
}

func NewCheckTicketMiddleware(secret string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ticket := r.URL.Query().Get("ticket")
			if ticket == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			token, err := jwt.ParseWithClaims(ticket, &api.TicketClaims{}, func(token *jwt.Token) (any, error) {
				return []byte(secret), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(api.ErrorResponse{
					Message: "malformed accessToken",
				})
				return
			}

			claims, ok := token.Claims.(*api.TicketClaims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(api.ErrorResponse{
					Message: "unknown claims type, cannot proceed",
				})
				return
			}

			// TODO: check exp

			ctx := context.WithValue(r.Context(), "userID", claims.UserID)

			next(w, r.WithContext(ctx))
		}
	}
}
