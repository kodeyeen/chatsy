package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kodeyeen/chatsy/internal/api"
)

type handler struct {
	svc service
}

func NewHandler(s service) *handler {
	return &handler{
		svc: s,
	}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	var regData RegisterData
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(api.Error{
			Message: "invalid input data",
		})
		return
	}

	userDTO, err := h.svc.Register(r.Context(), &regData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(api.Error{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userDTO)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(api.Error{
			Message: "invalid input data",
		})
		return
	}

	loginResult, err := h.svc.Login(r.Context(), &creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(api.Error{
			Message: "wrong credentials",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    loginResult.AccessToken,
		Path:     "/",
		Expires:  loginResult.Exp,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	json.NewEncoder(w).Encode(loginResult.User)
}

func (h *handler) Logout(w http.ResponseWriter, r *http.Request) {
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

func (h *handler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userDTO, err := h.svc.GetUserByID(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// json.NewEncoder(w).Encode(api.Error{
		// 	Message: "invalid accessToken was provided",
		// })
		return
	}

	json.NewEncoder(w).Encode(userDTO)
}

func (h *handler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ticketDTO := h.svc.CreateTicket(ctx, userID)

	json.NewEncoder(w).Encode(ticketDTO)
}
