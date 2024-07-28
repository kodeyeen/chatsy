package restapi

import "time"

type GetUserResponse struct {
	ID           *int       `json:"id"`
	Username     *string    `json:"username"`
	FirstName    *string    `json:"firstName"`
	LastName     *string    `json:"lastName"`
	Email        *string    `json:"email"`
	PasswordHash *string    `json:"-"`
	JoinedAt     *time.Time `json:"joinedAt"`
}
