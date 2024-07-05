package user

import "context"

type repository interface {
	Add(ctx context.Context, u *User) error
	FindByID(ctx context.Context, id int) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type service struct {
	repo repository
}

func NewDefaultService(repo repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetByID(ctx context.Context, id int) (*GetResponse, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return &GetResponse{}, err
	}

	userDTO := GetResponse{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return &userDTO, nil
}
