package message

import (
	"context"
)

type repository interface {
	Add(context.Context, *Message) error
	FindForChat(ctx context.Context, chatID int, limit, offset int) ([]*Message, error)
	CountForChat(ctx context.Context, chatID int) (int, error)
}
