package chat

import (
	"context"

	"github.com/kodeyeen/chatsy/internal/entity"
	"github.com/kodeyeen/chatsy/v1"
)

type Repository interface {
	Add(context.Context, *entity.Chat) error
	FindByID(context.Context, int) (*entity.Chat, error)
	FindByUserID(ctx context.Context, userID int, limit, offset int) ([]*entity.Chat, error)
	FindAllForUser(ctx context.Context, userID int) ([]*entity.Chat, error)
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

func (s *Service) GetByID(ctx context.Context, id int) (*chatsy.GetChatResponse, error) {
	c, err := s.chats.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &chatsy.GetChatResponse{
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
		resp.LastMessage = &chatsy.GetMessageResponse{
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

func (s *Service) GetAllForUser(ctx context.Context, userID int) ([]*chatsy.GetChatResponse, error) {
	cs, err := s.chats.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := make([]*chatsy.GetChatResponse, 0, len(cs))

	for _, c := range cs {
		resp = append(resp, &chatsy.GetChatResponse{
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

func (s *Service) GetByUserID(ctx context.Context, userID int, limit, offset int) (*chatsy.PageResponse[*chatsy.GetChatResponse], error) {
	cs, err := s.chats.FindByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	cnt, err := s.chats.CountForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	items := make([]*chatsy.GetChatResponse, 0, len(cs))

	for _, c := range cs {
		item := &chatsy.GetChatResponse{
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
			item.LastMessage = &chatsy.GetMessageResponse{
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

	resp := &chatsy.PageResponse[*chatsy.GetChatResponse]{
		Items:  items,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return resp, nil
}
