package user

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	dbpool *pgxpool.Pool
}

func NewPostgresRepository(dbpool *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{
		dbpool: dbpool,
	}
}

func (r *postgresRepository) Add(ctx context.Context, u *User) (int, error) {
	query := `
		INSERT INTO users (first_name, last_name, username, email, password_hash)
		VALUES (@first_name, @last_name, @username, @email, @password_hash)
		RETURNING id
	`
	args := pgx.NamedArgs{
		"first_name":    u.FirstName,
		"last_name":     u.LastName,
		"username":      u.Username,
		"email":         u.Email,
		"password_hash": u.PasswordHash,
	}

	var lastInsertId int
	err := r.dbpool.QueryRow(ctx, query, args).Scan(&lastInsertId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				if pgErr.ConstraintName == "users_username_key" {
					return 0, ErrUsernameAlreadyExists
				} else if pgErr.ConstraintName == "users_email_key" {
					return 0, ErrEmailAlreadyExists
				}
			}
		}

		return 0, err
	}

	return lastInsertId, nil
}

func (r *postgresRepository) FindByID(ctx context.Context, id int) (*User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, password_hash, joined_at
		FROM users
		WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": id,
	}

	rows, _ := r.dbpool.Query(ctx, query, args)
	userDTO, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[GetDTO])
	if err != nil {
		return &User{}, err
	}

	u := &User{
		ID:           userDTO.ID,
		FirstName:    userDTO.FirstName,
		LastName:     userDTO.LastName,
		Username:     userDTO.Username,
		Email:        userDTO.Email,
		PasswordHash: userDTO.PasswordHash,
		JoinedAt:     userDTO.JoinedAt,
	}

	return u, nil
}

func (r *postgresRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, password_hash, joined_at
		FROM users
		WHERE email = @email
	`
	args := pgx.NamedArgs{
		"email": email,
	}

	rows, _ := r.dbpool.Query(ctx, query, args)
	userDTO, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[GetDTO])
	if err != nil {
		return &User{}, err
	}

	u := &User{
		ID:           userDTO.ID,
		FirstName:    userDTO.FirstName,
		LastName:     userDTO.LastName,
		Username:     userDTO.Username,
		Email:        userDTO.Email,
		PasswordHash: userDTO.PasswordHash,
		JoinedAt:     userDTO.JoinedAt,
	}

	return u, nil
}
