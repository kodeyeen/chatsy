package message

import (
	"context"
	"log"

	"github.com/kodeyeen/chatsy/internal/dto"
	"github.com/kodeyeen/chatsy/internal/user"
)

type userRepository interface {
	FindByID(ctx context.Context, ID int) (*user.User, error)
}

type defaultService struct {
	repo    repository
	usrRepo userRepository
}

func NewDefaultService(repo repository, usrRepo userRepository) *defaultService {
	return &defaultService{
		repo:    repo,
		usrRepo: usrRepo,
	}
}

func (s *defaultService) Create(ctx context.Context, createDTO *CreateDTO, senderID int) (*GetDTO, error) {
	sender, err := s.usrRepo.FindByID(ctx, senderID)
	if err != nil {
		return &GetDTO{}, err
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
		return &GetDTO{}, err
	}

	msgDTO := &GetDTO{
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

func (s *defaultService) GetForChat(ctx context.Context, chatID int, limit, offset int) (*dto.PageResponse[GetDTO], error) {
	msgs, err := s.repo.FindForChat(ctx, chatID, limit, offset)
	if err != nil {
		log.Println("HERE 1", err)
		return &dto.PageResponse[GetDTO]{}, err
	}

	cnt, err := s.repo.CountForChat(ctx, chatID)
	if err != nil {
		log.Println("HERE 2", ctx)
		return &dto.PageResponse[GetDTO]{}, err
	}

	dtos := make([]*GetDTO, 0, len(msgs))

	for _, msg := range msgs {
		getDTO := &GetDTO{
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

	page := &dto.PageResponse[GetDTO]{
		Items:  dtos,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return page, nil
}
