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

type Service struct {
	users repository
}

func NewService(users repository) *Service {
	return &Service{
		users: users,
	}
}

func (s *Service) GetByID(ctx context.Context, id int) (*GetResponse, error) {
	u, err := s.users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := GetResponse{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return &resp, nil
}
