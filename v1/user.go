package chatsy

import "time"

type GetUserResponse struct {
	ID        *int       `json:"id"`
	Username  *string    `json:"username"`
	FirstName *string    `json:"firstName"`
	LastName  *string    `json:"lastName"`
	Email     *string    `json:"email"`
	JoinedAt  *time.Time `json:"joinedAt"`
}
