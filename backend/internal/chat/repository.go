package chat

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("chat not found")
)

type repository interface {
	Add(context.Context, *Chat) error
	FindByID(context.Context, int) (*Chat, error)
	FindForUser(ctx context.Context, userID int, limit, offset int) ([]*Chat, error)
	FindAllForUser(ctx context.Context, userID int) ([]*Chat, error)
	CountForUser(ctx context.Context, userID int) (int, error)
}
