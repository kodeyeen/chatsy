package user

import (
	"context"

	"github.com/kodeyeen/chatsy/internal/entity"
	"github.com/kodeyeen/chatsy/v1"
)

type Repository interface {
	Add(ctx context.Context, u *entity.User) error
	FindByID(ctx context.Context, id int) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type Service struct {
	users Repository
}

func NewService(users Repository) *Service {
	return &Service{
		users: users,
	}
}

func (s *Service) GetByID(ctx context.Context, id int) (*chatsy.GetUserResponse, error) {
	u, err := s.users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := chatsy.GetUserResponse{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return &resp, nil
}
