package message

import (
	"context"

	"github.com/kodeyeen/chatsy/internal/api"
)

type service interface {
	Create(ctx context.Context) error
	GetForChat(ctx context.Context, chatID int, limit, offset int) (*api.Page[GetResponse], error)
}
