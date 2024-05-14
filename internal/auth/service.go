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
	Register(ctx context.Context, regData *RegistrationRequest) (*user.GetResponse, error)
	Login(ctx context.Context, creds *Credentials) (*LoginResult, error)
	GetUserByID(ctx context.Context, id int) (*user.GetResponse, error)
	CreateTicket(ctx context.Context, userID int) (string, error)
}
