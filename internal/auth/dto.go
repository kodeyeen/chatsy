package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodeyeen/chatsy/internal/security"
	"github.com/kodeyeen/chatsy/internal/user"
)

type RegistrationRequest struct {
	Username  *string            `json:"username"`
	FirstName *string            `json:"firstName"`
	LastName  *string            `json:"lastName"`
	Email     *string            `json:"email"`
	Password  *security.Password `json:"password"`
}

type Credentials struct {
	Email    *string            `json:"email"`
	Password *security.Password `json:"password"`
}

type LoginResult struct {
	AccessToken *string `json:"accessToken"`
	Exp         *time.Time
	User        *user.GetDTO
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
