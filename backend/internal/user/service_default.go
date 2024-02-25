package user

import "context"

type defaultService struct {
	repo repository
}

func NewDefaultService(repo repository) *defaultService {
	return &defaultService{
		repo: repo,
	}
}

func (s *defaultService) GetByID(ctx context.Context, id int) (*GetDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return &GetDTO{}, err
	}

	userDTO := GetDTO{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return &userDTO, nil
}
