package chat

import (
	"context"

	"github.com/jackc/pgx/v5"
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

func (r *postgresRepository) Add(ctx context.Context, c *Chat) (int, error) {
	query := `
		INSERT INTO chats (type, title, invite_hash)
		VALUES (@type, @title, @invite_hash)
		RETURNING id
	`
	args := pgx.NamedArgs{
		"type":        c.Type,
		"title":       c.Title,
		"invite_hash": c.InviteHash,
	}

	var lastInsertId int
	err := r.dbpool.QueryRow(ctx, query, args).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (r *postgresRepository) FindByID(ctx context.Context, id int) (*Chat, error) {
	query := `
		SELECT
			id,
			type,
			title,
			description,
			invite_hash,
			join_by_link_count,
			is_private,
			join_by_request
		FROM chats
		WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": id,
	}

	rows, _ := r.dbpool.Query(ctx, query, args)
	dto, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[GetDTO])
	if err != nil {
		return &Chat{}, err
	}

	c := &Chat{
		ID:              dto.ID,
		Type:            dto.Type,
		Title:           dto.Title,
		Description:     dto.Description,
		InviteHash:      dto.InviteHash,
		JoinByLinkCount: dto.JoinByLinkCount,
		IsPrivate:       dto.IsPrivate,
		JoinByRequest:   dto.JoinByRequest,
	}

	return c, nil
}

func (r *postgresRepository) FindUserChats(ctx context.Context, userID int, limit, offset int) ([]*Chat, error) {
	query := `
		SELECT
			c.id,
			c.type,
			c.title,
			c.description,
			c.invite_hash,
			c.join_by_link_count,
			c.is_private,
			c.join_by_request
		FROM chats c
		INNER JOIN participations p ON c.id = p.chat_id
		WHERE p.user_id = @user_id
	`
	args := pgx.NamedArgs{
		"user_id": userID,
	}

	rows, _ := r.dbpool.Query(ctx, query, args)
	dtos, err := pgx.CollectRows(rows, pgx.RowToStructByName[GetDTO])
	if err != nil {
		return []*Chat{}, err
	}

	var chats []*Chat

	for _, dto := range dtos {
		chats = append(chats, &Chat{
			ID:              dto.ID,
			Type:            dto.Type,
			Title:           dto.Title,
			Description:     dto.Description,
			InviteHash:      dto.InviteHash,
			JoinByLinkCount: dto.JoinByLinkCount,
			IsPrivate:       dto.IsPrivate,
			JoinByRequest:   dto.JoinByRequest,
		})
	}

	return chats, nil
}

func (r *postgresRepository) CountUserChats(ctx context.Context, userID int) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM chats c
		INNER JOIN participations p ON c.id = p.chat_id
		WHERE p.user_id = @user_id
	`
	args := pgx.NamedArgs{
		"user_id": userID,
	}

	rows, err := r.dbpool.Query(ctx, query, args)
	if err != nil {
		return 0, err
	}

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
