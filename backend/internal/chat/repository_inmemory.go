package chat

import (
	"context"
)

type inmemoryRepository struct {
	chats map[int]*Chat
}

func NewInmemoryRepository() *inmemoryRepository {
	return &inmemoryRepository{
		chats: make(map[int]*Chat),
	}
}

func (r *inmemoryRepository) Add(ctx context.Context, c *Chat) (int, error) {
	id := 5

	r.chats[id] = c

	return id, nil
}

func (r *inmemoryRepository) FindByID(ctx context.Context, id int) (*Chat, error) {
	return nil, nil
}
