package chat

import (
	"context"
)

type service interface {
	Create(ctx context.Context) error
	GetByID(ctx context.Context) error
	GetUserChats(ctx context.Context, userID int, limit, offset int) (*ChatPage, error)
}
