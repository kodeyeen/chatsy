package message

import (
	"context"
	"log"

	"github.com/kodeyeen/chatsy/internal/api"
	"github.com/kodeyeen/chatsy/internal/user"
)

type repository interface {
	Add(context.Context, *Message) error
	FindForChat(ctx context.Context, chatID int, limit, offset int) ([]*Message, error)
	CountForChat(ctx context.Context, chatID int) (int, error)
}

type userRepository interface {
	FindByID(ctx context.Context, ID int) (*user.User, error)
}

type service struct {
	repo    repository
	usrRepo userRepository
}

func NewDefaultService(repo repository, usrRepo userRepository) *service {
	return &service{
		repo:    repo,
		usrRepo: usrRepo,
	}
}

func (s *service) Create(ctx context.Context, createDTO *CreateDTO, senderID int) (*GetResponse, error) {
	sender, err := s.usrRepo.FindByID(ctx, senderID)
	if err != nil {
		return &GetResponse{}, err
	}

	msg := &Message{
		ChatID:     createDTO.ChatID,
		SenderID:   sender.ID,
		SenderName: sender.Name,
		AuthorName: createDTO.AuthorName,
		OriginalID: createDTO.OriginalID,
		ParentID:   createDTO.ParentID,
		Text:       createDTO.Text,
		IsViewed:   new(bool),
	}

	err = s.repo.Add(ctx, msg)
	if err != nil {
		return &GetResponse{}, err
	}

	msgDTO := &GetResponse{
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

	return msgDTO, nil
}

func (s *service) GetForChat(ctx context.Context, chatID int, limit, offset int) (*api.Page[GetResponse], error) {
	msgs, err := s.repo.FindForChat(ctx, chatID, limit, offset)
	if err != nil {
		log.Println("HERE 1", err)
		return &api.Page[GetResponse]{}, err
	}

	cnt, err := s.repo.CountForChat(ctx, chatID)
	if err != nil {
		log.Println("HERE 2", ctx)
		return &api.Page[GetResponse]{}, err
	}

	dtos := make([]*GetResponse, 0, len(msgs))

	for _, msg := range msgs {
		getDTO := &GetResponse{
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
		}

		dtos = append(dtos, getDTO)
	}

	page := &api.Page[GetResponse]{
		Items:  dtos,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return page, nil
}
