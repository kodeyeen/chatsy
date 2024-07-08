package user

import (
	"context"

	"github.com/kodeyeen/chatsy"
)

type repository interface {
	Add(ctx context.Context, u *chatsy.User) error
	FindByID(ctx context.Context, id int) (*chatsy.User, error)
	FindByEmail(ctx context.Context, email string) (*chatsy.User, error)
}

type DefaultService struct {
	repo repository
}

func NewDefaultService(repo repository) *DefaultService {
	return &DefaultService{
		repo: repo,
	}
}

func (s *DefaultService) GetByID(ctx context.Context, id int) (*Response, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return &Response{}, err
	}

	userDTO := Response{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return &userDTO, nil
}
