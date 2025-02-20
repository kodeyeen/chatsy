package domain

import "time"

type Message struct {
	ID               *int
	ChatID           *int
	SenderID         *int
	SenderName       *string
	AuthorName       *string
	OriginalID       *int
	ParentID         *int
	ParentSenderName *string
	ParentText       *string
	Text             *string
	IsViewed         *bool
	SentAt           *time.Time
}

type MessageView struct {
	MessageID *int
	UserID    *int
}
