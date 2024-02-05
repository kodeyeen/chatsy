package chat

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("chat not found")
)

type repository interface {
	Add(context.Context, *Chat) (int, error)
	FindByID(context.Context, int) (*Chat, error)
	FindUserChats(ctx context.Context, userID int, limit, offset int) ([]*Chat, error)
	CountUserChats(ctx context.Context, userID int) (int, error)
}
