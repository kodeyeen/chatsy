package user

import (
	"context"

	"github.com/kodeyeen/chatsy/api/v1"
	"github.com/kodeyeen/chatsy/internal/domain"
)

type Repository interface {
	Add(ctx context.Context, u *domain.User) error
	FindByID(ctx context.Context, id int) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type Service struct {
	users Repository
}

func NewService(users Repository) *Service {
	return &Service{
		users: users,
	}
}

func (s *Service) GetByID(ctx context.Context, id int) (*api.GetUserResponse, error) {
	u, err := s.users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := api.GetUserResponse{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return &resp, nil
}
