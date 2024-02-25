package auth

import (
	"context"
	"errors"

	"github.com/kodeyeen/chatsy/internal/user"
)

var (
	ErrWrongCreds = errors.New("wrong credentials")
)

type service interface {
	Register(ctx context.Context, regData *RegisterData) (*user.GetDTO, error)
	Login(ctx context.Context, creds *Credentials) (*LoginResult, error)
	GetUserByID(ctx context.Context, id int) (*user.GetDTO, error)
	CreateTicket(ctx context.Context, userID int) (string, error)
}
