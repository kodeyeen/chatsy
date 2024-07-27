package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodeyeen/chatsy/internal/user"
)

type RegisterRequest struct {
	Username  *string `json:"username"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
}

type LoginRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type LoginResponse struct {
	AccessToken *string `json:"accessToken"`
	Exp         *time.Time
	User        *user.GetResponse
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID    int    `json:"userID"`
	UserEmail string `json:"userEmail"`
	Exp       int64  `json:"exp"`
}

type TicketClaims struct {
	jwt.RegisteredClaims
	UserID int   `json:"userID"`
	Exp    int64 `json:"exp"`
}
