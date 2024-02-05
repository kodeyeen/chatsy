package message

import "time"

type Message struct {
	ChatID            int
	AuthorID          int
	ForwardOriginalID int
	ReplyToID         int
	Text              string
	SentAt            time.Time
}
