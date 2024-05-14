package chat

import (
	"context"

	"github.com/kodeyeen/chatsy/internal/api"
)

type service interface {
	Create(ctx context.Context) error
	GetByID(ctx context.Context) error
	GetForUser(ctx context.Context, userID int, limit, offset int) (*api.Page[GetResponse], error)
}
