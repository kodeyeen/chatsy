package message

import (
	"context"

	"github.com/kodeyeen/chatsy/api/v1"
	"github.com/kodeyeen/chatsy/internal/domain"
)

type Repository interface {
	Add(context.Context, *domain.Message) error
	FindForChat(ctx context.Context, chatID int, limit, offset int) ([]*domain.Message, error)
	CountForChat(ctx context.Context, chatID int) (int, error)
}

type UserRepository interface {
	FindByID(ctx context.Context, ID int) (*domain.User, error)
}

type Service struct {
	messages Repository
	users    UserRepository
}

func NewService(messages Repository, users UserRepository) *Service {
	return &Service{
		messages: messages,
		users:    users,
	}
}

func (s *Service) Create(ctx context.Context, req *api.CreateMessageRequest, senderID int) (*api.GetMessageResponse, error) {
	sender, err := s.users.FindByID(ctx, senderID)
	if err != nil {
		return nil, err
	}

	msg := &domain.Message{
		ChatID:     req.ChatID,
		SenderID:   sender.ID,
		SenderName: sender.Name,
		AuthorName: req.AuthorName,
		OriginalID: req.OriginalID,
		ParentID:   req.ParentID,
		Text:       req.Text,
		IsViewed:   new(bool),
	}

	err = s.messages.Add(ctx, msg)
	if err != nil {
		return nil, err
	}

	resp := &api.GetMessageResponse{
		ID:         msg.ID,
		ChatID:     msg.ChatID,
		SenderID:   msg.SenderID,
		SenderName: msg.SenderName,
		AuthorName: msg.AuthorName,
		OriginalID: msg.OriginalID,
		ParentID:   msg.ParentID,
		Text:       msg.Text,
		SentAt:     msg.SentAt,
		IsViewed:   msg.IsViewed,
	}

	return resp, nil
}

func (s *Service) GetForChat(ctx context.Context, chatID int, limit, offset int) (*api.PageResponse[*api.GetMessageResponse], error) {
	msgs, err := s.messages.FindForChat(ctx, chatID, limit, offset)
	if err != nil {
		return nil, err
	}

	cnt, err := s.messages.CountForChat(ctx, chatID)
	if err != nil {
		return nil, err
	}

	items := make([]*api.GetMessageResponse, 0, len(msgs))

	for _, msg := range msgs {
		items = append(items, &api.GetMessageResponse{
			ID:               msg.ID,
			ChatID:           msg.ChatID,
			SenderID:         msg.SenderID,
			SenderName:       msg.SenderName,
			AuthorName:       msg.AuthorName,
			OriginalID:       msg.OriginalID,
			ParentID:         msg.ParentID,
			ParentSenderName: msg.ParentSenderName,
			ParentText:       msg.ParentText,
			Text:             msg.Text,
			SentAt:           msg.SentAt,
			IsViewed:         msg.IsViewed,
		})
	}

	resp := &api.PageResponse[*api.GetMessageResponse]{
		Items:  items,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return resp, nil
}
