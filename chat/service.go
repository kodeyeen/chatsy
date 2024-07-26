package chat

import (
	"context"

	"github.com/kodeyeen/chatsy"
	"github.com/kodeyeen/chatsy/api"
	"github.com/kodeyeen/chatsy/message"
)

type Repository interface {
	Add(context.Context, *chatsy.Chat) error
	FindByID(context.Context, int) (*chatsy.Chat, error)
	FindForUser(ctx context.Context, userID int, limit, offset int) ([]*chatsy.Chat, error)
	FindAllForUser(ctx context.Context, userID int) ([]*chatsy.Chat, error)
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

func (s *Service) GetByID(ctx context.Context, id int) (*GetResponse, error) {
	c, err := s.chats.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	getDTO := &GetResponse{
		ID:              c.ID,
		Type:            c.Type,
		Title:           c.Title,
		Description:     c.Description,
		InviteHash:      c.InviteHash,
		JoinByLinkCount: c.JoinByLinkCount,
		IsPrivate:       c.IsPrivate,
		JoinByRequest:   c.JoinByRequest,
	}

	if c.LastMessage != nil {
		getDTO.LastMessage = &message.GetResponse{
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

	return getDTO, nil
}

func (s *Service) GetAllForUser(ctx context.Context, userID int) ([]*GetResponse, error) {
	cs, err := s.chats.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := make([]*GetResponse, 0, len(cs))

	for _, c := range cs {
		resp = append(resp, &GetResponse{
			ID:                      c.ID,
			Type:                    c.Type,
			Title:                   c.Title,
			Description:             c.Description,
			InviteHash:              c.InviteHash,
			JoinByLinkCount:         c.JoinByLinkCount,
			IsPrivate:               c.IsPrivate,
			IsJoined:                c.IsJoined,
			ParticipantCount:        c.ParticipantCount,
			AreNotificationsEnabled: c.AreNotificationsEnabled,
			JoinByRequest:           c.JoinByRequest,
		})
	}

	return resp, nil
}

func (s *Service) GetForUser(ctx context.Context, userID int, limit, offset int) (*api.PageResponse[*GetResponse], error) {
	cs, err := s.chats.FindForUser(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	cnt, err := s.chats.CountForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	items := make([]*GetResponse, 0, len(cs))

	for _, c := range cs {
		item := &GetResponse{
			ID:                      c.ID,
			Type:                    c.Type,
			Title:                   c.Title,
			Description:             c.Description,
			InviteHash:              c.InviteHash,
			JoinByLinkCount:         c.JoinByLinkCount,
			IsPrivate:               c.IsPrivate,
			IsJoined:                c.IsJoined,
			ParticipantCount:        c.ParticipantCount,
			AreNotificationsEnabled: c.AreNotificationsEnabled,
			JoinByRequest:           c.JoinByRequest,
		}

		if c.LastMessage != nil {
			item.LastMessage = &message.GetResponse{
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

	resp := &api.PageResponse[*GetResponse]{
		Items:  items,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return resp, nil
}