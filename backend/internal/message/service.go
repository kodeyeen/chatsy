package message

import (
	"context"

	"github.com/kodeyeen/chatsy/internal/dto"
)

type service interface {
	Create(ctx context.Context) error
	GetForChat(ctx context.Context, chatID int, limit, offset int) (*dto.Page[GetDTO], error)
}
