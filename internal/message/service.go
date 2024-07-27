package message

import (
	"context"

	"github.com/kodeyeen/chatsy"
	"github.com/kodeyeen/chatsy/internal/api"
)

type Repository interface {
	Add(context.Context, *chatsy.Message) error
	FindForChat(ctx context.Context, chatID int, limit, offset int) ([]*chatsy.Message, error)
	CountForChat(ctx context.Context, chatID int) (int, error)
}

type UserRepository interface {
	FindByID(ctx context.Context, ID int) (*chatsy.User, error)
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

func (s *Service) Create(ctx context.Context, createDTO *CreateDTO, senderID int) (*GetResponse, error) {
	sender, err := s.users.FindByID(ctx, senderID)
	if err != nil {
		return nil, err
	}

	msg := &chatsy.Message{
		ChatID:     createDTO.ChatID,
		SenderID:   sender.ID,
		SenderName: sender.Name,
		AuthorName: createDTO.AuthorName,
		OriginalID: createDTO.OriginalID,
		ParentID:   createDTO.ParentID,
		Text:       createDTO.Text,
		IsViewed:   new(bool),
	}

	err = s.messages.Add(ctx, msg)
	if err != nil {
		return nil, err
	}

	resp := &GetResponse{
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

func (s *Service) GetForChat(ctx context.Context, chatID int, limit, offset int) (*api.PageResponse[*GetResponse], error) {
	msgs, err := s.messages.FindForChat(ctx, chatID, limit, offset)
	if err != nil {
		return nil, err
	}

	cnt, err := s.messages.CountForChat(ctx, chatID)
	if err != nil {
		return nil, err
	}

	items := make([]*GetResponse, 0, len(msgs))

	for _, msg := range msgs {
		items = append(items, &GetResponse{
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

	resp := &api.PageResponse[*GetResponse]{
		Items:  items,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return resp, nil
}