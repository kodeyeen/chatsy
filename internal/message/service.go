package message

import (
	"context"

	"github.com/kodeyeen/chatsy/internal/domain"
	"github.com/kodeyeen/chatsy/v1"
)

type Repository interface {
	Add(context.Context, *domain.Message) error
	FindByChatID(ctx context.Context, chatID int, limit, offset int) ([]*domain.Message, error)
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

func (s *Service) Create(ctx context.Context, req *chatsy.CreateMessageRequest, senderID int) (*chatsy.GetMessageResponse, error) {
	sender, err := s.users.FindByID(ctx, senderID)
	if err != nil {
		return nil, err
	}

	tmp := "tmp"

	msg := &domain.Message{
		ChatID:     req.ChatID,
		SenderID:   sender.ID,
		SenderName: &tmp,
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

	resp := &chatsy.GetMessageResponse{
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

func (s *Service) GetByChatID(ctx context.Context, chatID int, limit, offset int) (*chatsy.PageResponse[*chatsy.GetMessageResponse], error) {
	msgs, err := s.messages.FindByChatID(ctx, chatID, limit, offset)
	if err != nil {
		return nil, err
	}

	cnt, err := s.messages.CountForChat(ctx, chatID)
	if err != nil {
		return nil, err
	}

	items := make([]*chatsy.GetMessageResponse, 0, len(msgs))

	for _, msg := range msgs {
		items = append(items, &chatsy.GetMessageResponse{
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

	resp := &chatsy.PageResponse[*chatsy.GetMessageResponse]{
		Items:  items,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return resp, nil
}
