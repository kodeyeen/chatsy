package user

import (
	"context"
)

type defaultService struct {
	repo repository
}

func NewDefaultService(repo repository) *defaultService {
	return &defaultService{
		repo: repo,
	}
}

func (s *defaultService) GetCurrentUser(ctx context.Context) error {
	return nil
}
