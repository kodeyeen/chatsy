package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodeyeen/chatsy"
	"github.com/kodeyeen/chatsy/user"
)

var (
	ErrWrongCreds = errors.New("wrong credentials")
)

type UserRepository interface {
	Add(ctx context.Context, usr *chatsy.User) error
	FindByID(ctx context.Context, id int) (*chatsy.User, error)
	FindByEmail(ctx context.Context, email string) (*chatsy.User, error)
}

type Service struct {
	users     UserRepository
	secret    string
	tokenTTL  time.Duration
	ticketTTL time.Duration
}

func NewService(
	secret string,
	tokenTTL time.Duration,
	ticketTTL time.Duration,
	users UserRepository,
) *Service {
	return &Service{
		secret:    secret,
		tokenTTL:  tokenTTL,
		ticketTTL: ticketTTL,
		users:     users,
	}
}

func (s *Service) Register(ctx context.Context, regData *RegisterRequest) (*user.GetResponse, error) {
	passwordHash, err := regData.Password.Hash()
	if err != nil {
		return nil, err
	}

	usr := &chatsy.User{
		FirstName:    regData.FirstName,
		LastName:     regData.LastName,
		Username:     regData.Username,
		Email:        regData.Email,
		PasswordHash: &passwordHash,
	}

	err = s.users.Add(ctx, usr)
	if err != nil {
		return nil, err
	}

	resp := &user.GetResponse{
		ID:        usr.ID,
		Username:  usr.Username,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		JoinedAt:  usr.JoinedAt,
	}

	return resp, nil
}

func (s *Service) Login(ctx context.Context, creds *LoginRequest) (*LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	usr, err := s.users.FindByEmail(ctx, *creds.Email)
	if err != nil {
		return nil, ErrWrongCreds
	}

	err = creds.Password.Matches(*usr.PasswordHash)
	if err != nil {
		return nil, ErrWrongCreds
	}

	exp := time.Now().Add(s.tokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		UserID:    *usr.ID,
		UserEmail: *usr.Email,
		Exp:       exp.Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	usrResp := user.GetResponse{
		ID:        usr.ID,
		Username:  usr.Username,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		JoinedAt:  usr.JoinedAt,
	}

	return &LoginResponse{
		AccessToken: &tokenString,
		Exp:         &exp,
		User:        &usrResp,
	}, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*user.GetResponse, error) {
	usr, err := s.users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := user.GetResponse{
		ID:        usr.ID,
		Username:  usr.Username,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		JoinedAt:  usr.JoinedAt,
	}

	return &resp, nil
}

func (s *Service) CreateTicket(ctx context.Context, userID int) (string, error) {
	exp := time.Now().Add(s.ticketTTL)
	claims := TicketClaims{
		UserID: userID,
		Exp:    exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}