package postgres

import (
	"context"
	"errors"

	"github.com/kodeyeen/chatsy/internal/database"
	"github.com/kodeyeen/chatsy/internal/entity"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	dbpool *pgxpool.Pool
}

func NewUserRepository(dbpool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		dbpool: dbpool,
	}
}

func (r *UserRepository) Add(ctx context.Context, usr *entity.User) error {
	query := `
		INSERT INTO users (first_name, last_name, username, email, password_hash)
		VALUES (@first_name, @last_name, @username, @email, @password_hash)
		RETURNING id
	`
	args := pgx.NamedArgs{
		"first_name":    usr.FirstName,
		"last_name":     usr.LastName,
		"username":      usr.Username,
		"email":         usr.Email,
		"password_hash": usr.PasswordHash,
	}

	err := r.dbpool.QueryRow(ctx, query, args).Scan(&usr.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				if pgErr.ConstraintName == "users_username_key" {
					return database.ErrUsernameAlreadyExists
				} else if pgErr.ConstraintName == "users_email_key" {
					return database.ErrEmailAlreadyExists
				}
			}
		}

		return err
	}

	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, ID int) (*entity.User, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			username,
			email,
			password_hash,
			joined_at
		FROM users
		WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": ID,
	}

	var usr entity.User
	err := r.dbpool.QueryRow(ctx, query, args).Scan(
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.Username,
		&usr.Email,
		&usr.PasswordHash,
		&usr.JoinedAt,
	)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, password_hash, joined_at
		FROM users
		WHERE email = @email
	`
	args := pgx.NamedArgs{
		"email": email,
	}

	var usr entity.User
	err := r.dbpool.QueryRow(ctx, query, args).Scan(
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.Username,
		&usr.Email,
		&usr.PasswordHash,
		&usr.JoinedAt,
	)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
