package chat

import (
	"context"
)

type defaultService struct {
	repo repository
}

func NewDefaultService(repo repository) *defaultService {
	return &defaultService{
		repo: repo,
	}
}

func (s *defaultService) Create(ctx context.Context) error {
	return nil
}

func (s *defaultService) GetByID(ctx context.Context) error {
	return nil
}

func (s *defaultService) GetUserChats(ctx context.Context, userID int, limit, offset int) (*ChatPage, error) {
	cs, err := s.repo.FindUserChats(ctx, userID, limit, offset)
	if err != nil {
		return &ChatPage{}, err
	}

	cnt, err := s.repo.CountUserChats(ctx, userID)
	if err != nil {
		return &ChatPage{}, err
	}

	var dtos []*GetDTO

	for _, c := range cs {
		dtos = append(dtos, &GetDTO{
			ID:              c.ID,
			Type:            c.Type,
			Title:           c.Title,
			Description:     c.Description,
			InviteHash:      c.InviteHash,
			JoinByLinkCount: c.JoinByLinkCount,
			IsPrivate:       c.IsPrivate,
			JoinByRequest:   c.JoinByRequest,
		})
	}

	page := &ChatPage{
		Items:  dtos,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return page, nil
}
