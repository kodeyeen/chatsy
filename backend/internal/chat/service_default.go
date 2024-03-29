package chat

import (
	"context"

	"github.com/kodeyeen/chatsy/internal/dto"
	"github.com/kodeyeen/chatsy/internal/message"
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

func (s *defaultService) GetByID(ctx context.Context, id int) (*GetDTO, error) {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return &GetDTO{}, err
	}

	getDTO := &GetDTO{
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
		getDTO.LastMessage = &message.GetDTO{
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

func (s *defaultService) GetAllForUser(ctx context.Context, userID int) ([]*GetDTO, error) {
	cs, err := s.repo.FindAllForUser(ctx, userID)
	if err != nil {
		return []*GetDTO{}, err
	}

	dtos := make([]*GetDTO, 0, len(cs))

	for _, c := range cs {
		getDTO := &GetDTO{
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

		dtos = append(dtos, getDTO)
	}

	return dtos, nil
}

func (s *defaultService) GetForUser(ctx context.Context, userID int, limit, offset int) (*dto.Page[GetDTO], error) {
	cs, err := s.repo.FindForUser(ctx, userID, limit, offset)
	if err != nil {
		return &dto.Page[GetDTO]{}, err
	}

	cnt, err := s.repo.CountForUser(ctx, userID)
	if err != nil {
		return &dto.Page[GetDTO]{}, err
	}

	dtos := make([]*GetDTO, 0, len(cs))

	for _, c := range cs {
		getDTO := &GetDTO{
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
			getDTO.LastMessage = &message.GetDTO{
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

		dtos = append(dtos, getDTO)
	}

	page := &dto.Page[GetDTO]{
		Items:  dtos,
		Count:  cnt,
		Limit:  limit,
		Offset: offset,
	}

	return page, nil
}
