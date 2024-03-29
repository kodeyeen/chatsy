package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kodeyeen/chatsy/internal/dto"
)

type httpHandler struct {
	svc service
}

func NewHTTPHandler(s service) *httpHandler {
	return &httpHandler{
		svc: s,
	}
}

func (h *httpHandler) Register(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	var regData RegisterData
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.APIError{
			Message: "invalid input data",
		})
		return
	}

	userDTO, err := h.svc.Register(r.Context(), &regData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.APIError{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userDTO)
}

func (h *httpHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.APIError{
			Message: "invalid input data",
		})
		return
	}

	loginResult, err := h.svc.Login(r.Context(), &creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(dto.APIError{
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

func (h *httpHandler) Logout(w http.ResponseWriter, r *http.Request) {
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

func (h *httpHandler) Me(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userDTO, err := h.svc.GetUserByID(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// json.NewEncoder(w).Encode(dto.APIError{
		// 	Message: "invalid accessToken was provided",
		// })
		return
	}

	json.NewEncoder(w).Encode(userDTO)
}

func (h *httpHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ticket, err := h.svc.CreateTicket(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ticket)
}
