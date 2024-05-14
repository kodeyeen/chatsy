package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodeyeen/chatsy/internal/user"
)

type userRepository interface {
	Add(ctx context.Context, usr *user.User) error
	FindByID(ctx context.Context, id int) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}

type defaultService struct {
	userRepo  userRepository
	secret    string
	tokenTTL  time.Duration
	ticketTTL time.Duration
}

func NewDefaultService(
	secret string,
	tokenTTL time.Duration,
	ticketTTL time.Duration,
	userRepo userRepository,
) *defaultService {
	return &defaultService{
		secret:    secret,
		tokenTTL:  tokenTTL,
		ticketTTL: ticketTTL,
		userRepo:  userRepo,
	}
}

func (s *defaultService) Register(ctx context.Context, regData *RegistrationRequest) (*user.Response, error) {
	passwordHash, err := regData.Password.Hash()
	if err != nil {
		return &user.Response{}, err
	}

	usr := &user.User{
		FirstName:    regData.FirstName,
		LastName:     regData.LastName,
		Username:     regData.Username,
		Email:        regData.Email,
		PasswordHash: &passwordHash,
	}

	err = s.userRepo.Add(ctx, usr)
	if err != nil {
		return &user.Response{}, err
	}

	userDTO := &user.Response{
		ID:        usr.ID,
		Username:  usr.Username,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		JoinedAt:  usr.JoinedAt,
	}

	return userDTO, nil
}

func (s *defaultService) Login(ctx context.Context, creds *Credentials) (*LoginResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	usr, err := s.userRepo.FindByEmail(ctx, *creds.Email)
	if err != nil {
		return &LoginResult{}, ErrWrongCreds
	}

	err = creds.Password.Matches(*usr.PasswordHash)
	if err != nil {
		return &LoginResult{}, ErrWrongCreds
	}

	exp := time.Now().Add(s.tokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		UserID:    *usr.ID,
		UserEmail: *usr.Email,
		Exp:       exp.Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return &LoginResult{}, err
	}

	userDTO := user.Response{
		ID:        usr.ID,
		Username:  usr.Username,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		JoinedAt:  usr.JoinedAt,
	}

	return &LoginResult{
		AccessToken: &tokenString,
		Exp:         &exp,
		User:        &userDTO,
	}, nil
}

func (s *defaultService) GetUserByID(ctx context.Context, id int) (*user.Response, error) {
	usr, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return &user.Response{}, err
	}

	userDTO := user.Response{
		ID:        usr.ID,
		Username:  usr.Username,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		JoinedAt:  usr.JoinedAt,
	}

	return &userDTO, nil
}

func (s *defaultService) CreateTicket(ctx context.Context, userID int) (string, error) {
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
