package chat

import (
	"context"

	"github.com/kodeyeen/chatsy/api/v1"
	"github.com/kodeyeen/chatsy/internal/domain"
)

type Repository interface {
	Add(context.Context, *domain.Chat) error
	FindByID(context.Context, int) (*domain.Chat, error)
	FindByUserID(ctx context.Context, userID int, limit, offset int) ([]*domain.Chat, error)
	FindAllForUser(ctx context.Context, userID int) ([]*domain.Chat, error)
	CountForUser(ctx context.Context, userID int) (int, error)
}

type Service struct {
	chats Repository
}

func NewService(chats Repository) *Service {
	return &Service{
		chats: chats,
	}
}

func (s *Service) Create(ctx context.Context) error {
	return nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*api.GetChatResponse, error) {
	c, err := s.chats.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &api.GetChatResponse{
		ID:              c.ID,
		Type:            (*string)(c.Type),
		Title:           c.Title,
		Description:     c.Description,
		InviteHash:      c.InviteHash,
		JoinByLinkCount: c.JoinByLinkCount,
		IsPrivate:       c.IsPrivate,
		JoinByRequest:   c.JoinByRequest,
	}

	if c.LastMessage != nil {
		resp.LastMessage = &api.GetMessageResponse{
			ID:         c.LastMessage.ID,
			ChatID:     c.LastMessage.ChatID,
			SenderID:   c.LastMessage.SenderID,
			SenderName: c.LastMessage.SenderName,
			AuthorName: c.LastMessage.AuthorName,
			OriginalID: c.LastMessage.OriginalID,
			ParentID:   c.LastMessage.ParentID,
			Text:       c.LastMessage.Text,
			SentAt:     c.LastMessage.SentAt,
			IsViewed:   c.LastMessage.IsViewed,
		}
	}

	return resp, nil
}

func (s *Service) GetAllForUser(ctx context.Context, userID int) ([]*api.GetChatResponse, error) {
	cs, err := s.chats.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := make([]*api.GetChatResponse, 0, len(cs))

	for _, c := range cs {
		resp = append(resp, &api.GetChatResponse{
			ID:                      c.ID,
			Type:                    (*string)(c.Type),
			Title:                   c.Title,
			Description:             c.Description,
			InviteHash:              c.InviteHash,
			JoinByLinkCount:         c.JoinByLinkCount,
			IsPrivate:               c.IsPrivate,
			IsJoined:                c.IsJoined,
			MemberCount:             c.MemberCount,
			AreNotificationsEnabled: c.AreNotificationsEnabled,
			JoinByRequest:           c.JoinByRequest,
		})
	}

	return resp, nil
}

func (s *Service) GetByUserID(ctx context.Context, userID int, limit, offset int) (*api.PageResponse[*api.GetChatResponse], error) {
	cs, err := s.chats.FindByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	cnt, err := s.chats.CountForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	items := make([]*api.GetChatResponse, 0, len(cs))

	for _, c := range cs {
		item := &api.GetChatResponse{
			ID:                      c.ID,
			Type:                    (*string)(c.Type),
			Title:                   c.Title,
			Description:             c.Description,
			InviteHash:              c.InviteHash,
			JoinByLinkCount:         c.JoinByLinkCount,
			IsPrivate:               c.IsPrivate,
			IsJoined:                c.IsJoined,
			MemberCount:             c.MemberCount,
			AreNotificationsEnabled: c.AreNotificationsEnabled,
			JoinByRequest:           c.JoinByRequest,
		}

		if c.LastMessage != nil {
			item.LastMessage = &api.GetMessageResponse{
				ID:         c.LastMessage.ID,
				ChatID:     c.LastMessage.ChatID,
				SenderID:   c.LastMessage.SenderID,
				SenderName: c.LastMessage.SenderName,
				AuthorName: c.LastMessage.AuthorName,
				OriginalID: c.LastMessage.OriginalID,
				ParentID:   c.LastMessage.ParentID,
				Text:       c.LastMessage.Text,
				SentAt:     c.LastMessage.SentAt,
				IsViewed:   c.LastMessage.IsViewed,
			}
		}

		items = append(items, item)
	}

	resp := &api.PageResponse[*api.GetChatResponse]{
		Items:  items,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return resp, nil
}
