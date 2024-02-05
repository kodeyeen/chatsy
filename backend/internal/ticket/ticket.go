package ticket

import (
	"time"
)

type Ticket struct {
	ID        string
	UserID    int
	CreatedAt time.Time
}
