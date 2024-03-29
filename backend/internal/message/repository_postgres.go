package message

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

func (r *postgresRepository) Add(ctx context.Context, msg *Message) error {
	query := `
		INSERT INTO
			messages (chat_id, sender_id, original_id, parent_id, text)
		VALUES
			(@chat_id, @sender_id, @original_id, @parent_id, @text)
		RETURNING id, sent_at
	`
	args := pgx.NamedArgs{
		"chat_id":     msg.ChatID,
		"sender_id":   msg.SenderID,
		"original_id": msg.OriginalID,
		"parent_id":   msg.ParentID,
		"text":        msg.Text,
	}

	err := r.dbpool.QueryRow(ctx, query, args).Scan(&msg.ID, &msg.SentAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) FindForChat(ctx context.Context, chatID int, limit, offset int) ([]*Message, error) {
	query := `
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
			NULLIF(
				TRIM(CONCAT(pu.first_name, ' ', pu.last_name)), ''
			) AS parent_sender_name,
			p.text AS parent_text,
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
		LEFT OUTER JOIN
			messages p ON m.parent_id = p.id
		LEFT OUTER JOIN
			users pu ON p.sender_id = pu.id
		WHERE
			m.chat_id = @chat_id
		ORDER BY
			m.sent_at DESC
		LIMIT @limit OFFSET @offset
	`
	args := pgx.NamedArgs{
		"chat_id": chatID,
		"limit":   limit,
		"offset":  offset,
	}

	rows, _ := r.dbpool.Query(ctx, query, args)

	var msgs []*Message

	for rows.Next() {
		var msg Message

		err := rows.Scan(
			&msg.ID,
			&msg.ChatID,
			&msg.SenderID,
			&msg.SenderName,
			&msg.AuthorName,
			&msg.OriginalID,
			&msg.ParentID,
			&msg.ParentSenderName,
			&msg.ParentText,
			&msg.Text,
			&msg.SentAt,
			&msg.IsViewed,
		)
		if err != nil {
			return []*Message{}, err
		}

		msgs = append(msgs, &msg)
	}

	return msgs, nil
}

func (r *postgresRepository) CountForChat(ctx context.Context, chatID int) (int, error) {
	query := `
		SELECT
			COUNT(*)
		FROM
			messages m
		WHERE
			m.chat_id = @chat_id
	`
	args := pgx.NamedArgs{
		"chat_id": chatID,
	}

	var cnt int
	err := r.dbpool.QueryRow(ctx, query, args).Scan(&cnt)
	if err != nil {
		return 0, err
	}

	return cnt, nil
}
