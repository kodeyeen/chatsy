package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/kodeyeen/chatsy/restapi"
)

type authService interface {
	Register(ctx context.Context, regData *restapi.RegisterRequest) (*restapi.GetUserResponse, error)
	Login(ctx context.Context, creds *restapi.LoginRequest) (*restapi.LoginResponse, error)
	GetUserByID(ctx context.Context, id int) (*restapi.GetUserResponse, error)
	CreateTicket(ctx context.Context, userID int) (string, error)
}

type AuthController struct {
	service authService
}

func NewAuthController(svc authService) *AuthController {
	return &AuthController{
		service: svc,
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	var regData restapi.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(restapi.ErrorResponse{
			Message: "invalid input data",
		})
		return
	}

	userResp, err := c.service.Register(r.Context(), &regData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(restapi.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResp)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var creds restapi.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(restapi.ErrorResponse{
			Message: "invalid input data",
		})
		return
	}

	loginResult, err := c.service.Login(r.Context(), &creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(restapi.ErrorResponse{
			Message: "wrong credentials",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    *loginResult.AccessToken,
		Path:     "/",
		Expires:  *loginResult.Exp,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	json.NewEncoder(w).Encode(loginResult.User)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
}

func (c *AuthController) Me(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userResp, err := c.service.GetUserByID(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// json.NewEncoder(w).Encode(restapi.ErrorResponse{
		// 	Message: "invalid accessToken was provided",
		// })
		return
	}

	json.NewEncoder(w).Encode(userResp)
}

func (c *AuthController) CreateTicket(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ticket, err := c.service.CreateTicket(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ticket)
}
