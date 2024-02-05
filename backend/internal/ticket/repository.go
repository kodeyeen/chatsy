package ticket

import "errors"

var (
	ErrNotFound = errors.New("ticket not found")
)

type repository interface {
	Add(t *Ticket) string
	FindByID(id string) (*Ticket, error)
}
