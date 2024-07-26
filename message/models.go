package message

import (
	"time"
)

type GetResponse struct {
	ID               *int       `json:"id"`
	ChatID           *int       `json:"chatId"`
	SenderID         *int       `json:"senderId"`
	SenderName       *string    `json:"senderName"`
	AuthorName       *string    `json:"authorName"`
	OriginalID       *int       `json:"originalId"`
	ParentID         *int       `json:"parentId"`
	ParentSenderName *string    `json:"parentSenderName"`
	ParentText       *string    `json:"parentText"`
	Text             *string    `json:"text"`
	SentAt           *time.Time `json:"sentAt"`
	IsViewed         *bool      `json:"isViewed"`
}

type CreateDTO struct {
	ChatID     *int    `json:"chatId"`
	ParentID   *int    `json:"parentId"`
	OriginalID *int    `json:"originalId"`
	AuthorName *string `json:"authorName"`
	Text       *string `json:"text"`
}
