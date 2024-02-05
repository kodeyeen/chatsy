package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodeyeen/chatsy/internal/api"
)

func NewCheckJWTMiddleware(secret string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("accessToken")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(api.Error{
					Message: "accessToken was not provided",
				})
				return
			}

			token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (any, error) {
				return []byte(secret), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(api.Error{
					Message: "malformed accessToken",
				})
				return
			}

			claims, ok := token.Claims.(*Claims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(api.Error{
					Message: "unknown claims type, cannot proceed",
				})
				return
			}

			ctx := context.WithValue(r.Context(), "userID", claims.UserID)

			next(w, r.WithContext(ctx))
		}
	}
}
