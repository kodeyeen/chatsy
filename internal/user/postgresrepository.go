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

func (r *postgresRepository) Add(ctx context.Context, usr *User) error {
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
					return ErrUsernameAlreadyExists
				} else if pgErr.ConstraintName == "users_email_key" {
					return ErrEmailAlreadyExists
				}
			}
		}

		return err
	}

	return nil
}

func (r *postgresRepository) FindByID(ctx context.Context, ID int) (*User, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			CONCAT(first_name, ' ', last_name) AS name,
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

	var usr User
	err := r.dbpool.QueryRow(ctx, query, args).Scan(
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.Name,
		&usr.Username,
		&usr.Email,
		&usr.PasswordHash,
		&usr.JoinedAt,
	)
	if err != nil {
		return &User{}, err
	}

	return &usr, nil
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
	userDTO, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[Response])
	if err != nil {
		return &User{}, err
	}

	usr := &User{
		ID:           userDTO.ID,
		FirstName:    userDTO.FirstName,
		LastName:     userDTO.LastName,
		Username:     userDTO.Username,
		Email:        userDTO.Email,
		PasswordHash: userDTO.PasswordHash,
		JoinedAt:     userDTO.JoinedAt,
	}

	return usr, nil
}
