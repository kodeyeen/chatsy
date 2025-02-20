package chat

import (
	"context"

	"github.com/kodeyeen/chatsy/internal/domain"
	"github.com/kodeyeen/chatsy/v1"
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

func (s *Service) GetByID(ctx context.Context, req *chatsy.GetChatByIDRequest) (*chatsy.GetChatByIDResponse, error) {
	c, err := s.chats.FindByID(ctx, *req.ID)
	if err != nil {
		return nil, err
	}

	resp := &chatsy.GetChatByIDResponse{
		ChatResponse: &chatsy.ChatResponse{
			ID:              c.ID,
			Type:            (*string)(c.Type),
			Title:           c.Title,
			Description:     c.Description,
			InviteHash:      c.InviteHash,
			JoinByLinkCount: c.JoinByLinkCount,
			IsPrivate:       c.IsPrivate,
			JoinByRequest:   c.JoinByRequest,
		},
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

func (s *Service) GetAllForUser(ctx context.Context, req *chatsy.GetUserChatsRequest) ([]*chatsy.GetUserChatsResponse, error) {
	cs, err := s.chats.FindAllForUser(ctx, *req.UserID)
	if err != nil {
		return nil, err
	}

	resp := make([]*chatsy.GetUserChatsResponse, 0, len(cs))

	for _, c := range cs {
		resp = append(resp, &chatsy.GetUserChatsResponse{
			ChatResponse: &chatsy.ChatResponse{
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
			},
		})
	}

	return resp, nil
}

func (s *Service) GetByUserID(ctx context.Context, req *chatsy.GetUserChatsPageRequest) (*chatsy.GetUserChatsPageResponse, error) {
	cs, err := s.chats.FindByUserID(ctx, *req.UserID, req.Pagination.Limit, req.Pagination.Offset)
	if err != nil {
		return nil, err
	}

	cnt, err := s.chats.CountForUser(ctx, *req.UserID)
	if err != nil {
		return nil, err
	}

	items := make([]*chatsy.ChatResponse, 0, len(cs))

	for _, c := range cs {
		item := &chatsy.ChatResponse{
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

	resp := &chatsy.GetUserChatsPageResponse{
		PageResponse: &chatsy.PageResponse[*chatsy.ChatResponse]{
			Items:  items,
			Count:  cnt,
			Limit:  req.Pagination.Limit,
			Offset: req.Pagination.Offset,
		},
	}

	return resp, nil
}
