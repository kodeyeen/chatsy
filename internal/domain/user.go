package domain

import (
	"time"
)

type User struct {
	ID           *int
	FirstName    *string
	LastName     *string
	Name         *string
	Username     *string
	Email        *string
	PasswordHash *string
	JoinedAt     *time.Time
}
