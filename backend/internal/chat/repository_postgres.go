package chat

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodeyeen/chatsy/internal/message"
)

type postgresRepository struct {
	dbpool *pgxpool.Pool
}

func NewPostgresRepository(dbpool *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{
		dbpool: dbpool,
	}
}

func (r *postgresRepository) Add(ctx context.Context, c *Chat) error {
	query := `
		INSERT INTO
			chats (type, title, invite_hash)
		VALUES
			(@type, @title, @invite_hash)
		RETURNING id
	`
	args := pgx.NamedArgs{
		"type":        c.Type,
		"title":       c.Title,
		"invite_hash": c.InviteHash,
	}

	err := r.dbpool.QueryRow(ctx, query, args).Scan(&c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) FindByID(ctx context.Context, id int) (*Chat, error) {
	query := `
		SELECT
			c.id,
			c.type,
			c.title,
			c.description,
			c.invite_hash,
			c.join_by_link_count,
			c.is_private,
			c.join_by_request,
			m.*
		FROM
			chats c
		LEFT JOIN LATERAL (
			SELECT
				m.id,
				m.chat_id,
				m.sender_id,
				CONCAT(u.first_name, ' ', u.last_name) AS sender_name,
				CASE
					WHEN m.original_id IS NULL THEN NULL
					ELSE (
						SELECT CONCAT(ou.first_name, ' ', ou.last_name)
						FROM messages o
						LEFT OUTER JOIN users ou ON o.sender_id = ou.id
						WHERE o.id = m.original_id
					)
				END AS author_name,
				m.original_id,
				m.parent_id,
				m.text,
				m.sent_at,
				EXISTS (
					SELECT 1
					FROM message_views mv
					WHERE mv.message_id = m.id
				) AS is_viewed
			FROM
				messages m
			LEFT OUTER JOIN
				users u ON m.sender_id = u.id
			WHERE
				m.chat_id = c.id
			ORDER BY
				m.sent_at DESC
			LIMIT 1
		) m ON TRUE
		WHERE
			c.id = @id
		LIMIT 1
	`
	args := pgx.NamedArgs{
		"id": id,
	}

	var chat Chat
	var lastMsg message.Message
	err := r.dbpool.QueryRow(ctx, query, args).Scan(
		&chat.ID,
		&chat.Type,
		&chat.Title,
		&chat.Description,
		&chat.InviteHash,
		&chat.JoinByLinkCount,
		&chat.IsPrivate,
		&chat.JoinByRequest,

		&lastMsg.ID,
		&lastMsg.ChatID,
		&lastMsg.SenderID,
		&lastMsg.SenderName,
		&lastMsg.AuthorName,
		&lastMsg.OriginalID,
		&lastMsg.ParentID,
		&lastMsg.Text,
		&lastMsg.SentAt,
		&lastMsg.IsViewed,
	)
	if err != nil {
		return &Chat{}, err
	}

	if lastMsg.ID != nil {
		chat.LastMessage = &lastMsg
	}

	return &chat, nil
}

func (r *postgresRepository) FindAllForUser(ctx context.Context, userID int) ([]*Chat, error) {
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
		FROM
			chats c
		INNER JOIN
			participations p ON c.id = p.chat_id
		WHERE
			p.user_id = @user_id
	`
	args := pgx.NamedArgs{
		"user_id": userID,
	}

	rows, _ := r.dbpool.Query(ctx, query, args)

	var chats []*Chat

	for rows.Next() {
		var chat Chat

		err := rows.Scan(
			&chat.ID,
			&chat.Type,
			&chat.Title,
			&chat.Description,
			&chat.InviteHash,
			&chat.JoinByLinkCount,
			&chat.IsPrivate,
			&chat.JoinByRequest,
		)
		if err != nil {
			return []*Chat{}, err
		}

		chats = append(chats, &chat)
	}

	return chats, nil
}

func (r *postgresRepository) FindForUser(ctx context.Context, userID int, limit, offset int) ([]*Chat, error) {
	query := `
		SELECT
			c.id,
			c.type,
			c.title,
			c.description,
			c.invite_hash,
			c.join_by_link_count,
			c.is_private,
			c.join_by_request,
			TRUE as is_joined,
			(
				SELECT COUNT(*)
				FROM participations p
				WHERE p.chat_id = c.id
			) as participant_count,
			(
				SELECT p.are_notifications_enabled
				FROM participations p
				WHERE p.chat_id = c.id AND p.user_id = @user_id
			) as are_notifications_enabled,
			m.*
		FROM
			chats c
		INNER JOIN
			participations p ON c.id = p.chat_id
		LEFT JOIN LATERAL (
			SELECT
				m.id,
				m.chat_id,
				m.sender_id,
				CONCAT(u.first_name, ' ', u.last_name) AS sender_name,
				CASE
					WHEN m.original_id IS NULL THEN NULL
					ELSE (
						SELECT CONCAT(ou.first_name, ' ', ou.last_name)
						FROM messages o
						LEFT OUTER JOIN users ou ON o.sender_id = ou.id
						WHERE o.id = m.original_id
					)
				END AS author_name,
				m.original_id,
				m.parent_id,
				m.text,
				m.sent_at,
				EXISTS (
					SELECT 1
					FROM message_views mv
					WHERE mv.message_id = m.id
				) AS is_viewed
			FROM
				messages m
			LEFT OUTER JOIN
				users u ON m.sender_id = u.id
			WHERE
				m.chat_id = c.id
			ORDER BY
				m.sent_at DESC
			LIMIT 1
		) m ON true
		WHERE
			p.user_id = @user_id
		LIMIT @limit OFFSET @offset
	`
	args := pgx.NamedArgs{
		"user_id": userID,
		"limit":   limit,
		"offset":  offset,
	}

	rows, _ := r.dbpool.Query(ctx, query, args)

	var chats []*Chat

	for rows.Next() {
		var chat Chat
		var lastMsg message.Message

		err := rows.Scan(
			&chat.ID,
			&chat.Type,
			&chat.Title,
			&chat.Description,
			&chat.InviteHash,
			&chat.JoinByLinkCount,
			&chat.IsPrivate,
			&chat.JoinByRequest,
			&chat.IsJoined,
			&chat.ParticipantCount,
			&chat.AreNotificationsEnabled,

			&lastMsg.ID,
			&lastMsg.ChatID,
			&lastMsg.SenderID,
			&lastMsg.SenderName,
			&lastMsg.AuthorName,
			&lastMsg.OriginalID,
			&lastMsg.ParentID,
			&lastMsg.Text,
			&lastMsg.SentAt,
			&lastMsg.IsViewed,
		)
		if err != nil {
			return []*Chat{}, err
		}

		if lastMsg.ID != nil {
			chat.LastMessage = &lastMsg
		}

		chats = append(chats, &chat)
	}

	return chats, nil
}

func (r *postgresRepository) CountForUser(ctx context.Context, userID int) (int, error) {
	query := `
		SELECT
			COUNT(*)
		FROM
			participations p
		WHERE
			p.user_id = @user_id
	`
	args := pgx.NamedArgs{
		"user_id": userID,
	}

	var cnt int
	err := r.dbpool.QueryRow(ctx, query, args).Scan(&cnt)
	if err != nil {
		return 0, err
	}

	return cnt, nil
}
