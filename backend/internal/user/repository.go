package user

import (
	"context"
	"errors"
)

var (
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrNotFound              = errors.New("user not found")
)

type repository interface {
	Add(ctx context.Context, u *User) (int, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
