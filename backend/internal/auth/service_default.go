package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodeyeen/chatsy/internal/ticket"
	"github.com/kodeyeen/chatsy/internal/user"
)

type userRepository interface {
	Add(ctx context.Context, u *user.User) (int, error)
	FindByID(ctx context.Context, id int) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}

type ticketSaver interface {
	Add(ctx context.Context, t *ticket.Ticket) string
}

type defaultService struct {
	userRepo    userRepository
	secret      string
	tokenTTL    time.Duration
	ticketSaver ticketSaver
}

func NewDefaultService(
	secret string,
	tokenTTL time.Duration,
	userRepo userRepository,
	ticketSaver ticketSaver,
) *defaultService {
	return &defaultService{
		secret:      secret,
		tokenTTL:    tokenTTL,
		userRepo:    userRepo,
		ticketSaver: ticketSaver,
	}
}

func (s *defaultService) Register(ctx context.Context, regData *RegisterData) (*user.GetDTO, error) {
	passwordHash, err := regData.Password.Hash()
	if err != nil {
		return &user.GetDTO{}, err
	}

	u := &user.User{
		FirstName:    regData.FirstName,
		LastName:     regData.LastName,
		Username:     regData.Username,
		Email:        regData.Email,
		PasswordHash: passwordHash,
	}

	userID, err := s.userRepo.Add(ctx, u)
	if err != nil {
		return &user.GetDTO{}, err
	}

	u.ID = userID

	userDTO := &user.GetDTO{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return userDTO, nil
}

func (s *defaultService) Login(ctx context.Context, creds *Credentials) (*LoginResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	u, err := s.userRepo.FindByEmail(ctx, creds.Email)
	if err != nil {
		return &LoginResult{}, ErrWrongCreds
	}

	err = creds.Password.Matches(u.PasswordHash)
	if err != nil {
		return &LoginResult{}, ErrWrongCreds
	}

	exp := time.Now().Add(s.tokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID:    u.ID,
		UserEmail: u.Email,
		Exp:       exp.Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return &LoginResult{}, err
	}

	userDTO := user.GetDTO{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return &LoginResult{
		AccessToken: tokenString,
		Exp:         exp,
		User:        userDTO,
	}, nil
}

func (s *defaultService) GetUserByID(ctx context.Context, id int) (*user.GetDTO, error) {
	u, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return &user.GetDTO{}, err
	}

	userDTO := user.GetDTO{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		JoinedAt:  u.JoinedAt,
	}

	return &userDTO, nil
}

func (s *defaultService) CreateTicket(ctx context.Context, userID int) *TicketDTO {
	t := &ticket.Ticket{
		UserID: userID,
	}

	ticketID := s.ticketSaver.Add(ctx, t)

	t.ID = ticketID

	ticketDTO := &TicketDTO{
		ID: t.ID,
	}

	return ticketDTO
}
