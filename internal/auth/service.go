package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodeyeen/chatsy/internal/domain"
	"github.com/kodeyeen/chatsy/v1"
)

var (
	ErrWrongCreds = errors.New("wrong credentials")
)

type UserRepository interface {
	Add(ctx context.Context, usr *domain.User) error
	FindByID(ctx context.Context, id int) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
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

func (s *Service) Register(ctx context.Context, regData *chatsy.RegisterRequest) (*chatsy.GetUserResponse, error) {
	passwordHash, err := Password(*regData.Password).Hash()
	if err != nil {
		return nil, err
	}

	usr := &domain.User{
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

	resp := &chatsy.GetUserResponse{
		ID:        usr.ID,
		Username:  usr.Username,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		JoinedAt:  usr.JoinedAt,
	}

	return resp, nil
}

func (s *Service) Login(ctx context.Context, creds *chatsy.LoginRequest) (*chatsy.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	usr, err := s.users.FindByEmail(ctx, *creds.Email)
	fmt.Println("FOUND?", usr, err)
	if err != nil {
		return nil, ErrWrongCreds
	}

	err = Password(*creds.Password).Matches(*usr.PasswordHash)
	fmt.Println("MATCHES?", err)
	if err != nil {
		return nil, ErrWrongCreds
	}

	exp := time.Now().Add(s.tokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, chatsy.TokenClaims{
		UserID:    *usr.ID,
		UserEmail: *usr.Email,
		Exp:       exp.Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	usrResp := chatsy.GetUserResponse{
		ID:        usr.ID,
		Username:  usr.Username,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		JoinedAt:  usr.JoinedAt,
	}

	return &chatsy.LoginResponse{
		AccessToken: &tokenString,
		Exp:         &exp,
		User:        &usrResp,
	}, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*chatsy.GetUserResponse, error) {
	usr, err := s.users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := chatsy.GetUserResponse{
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
	claims := chatsy.TicketClaims{
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
