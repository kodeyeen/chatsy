package user

import (
	"time"
)

type Response struct {
	ID           *int       `json:"id" db:"id"`
	Username     *string    `json:"username" db:"username"`
	FirstName    *string    `json:"firstName" db:"first_name"`
	LastName     *string    `json:"lastName" db:"last_name"`
	Email        *string    `json:"email" db:"email"`
	PasswordHash *string    `json:"-" db:"password_hash"`
	JoinedAt     *time.Time `json:"joinedAt" db:"joined_at"`
}
