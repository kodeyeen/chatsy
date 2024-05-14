package chat

import "github.com/kodeyeen/chatsy/internal/message"

type Type string

const (
	TypeDialog Type = "dialog"
	TypeGroup  Type = "group"
)

type Chat struct {
	ID                      *int
	Type                    *Type
	Title                   *string
	Description             *string
	InviteHash              *string
	JoinByLinkCount         *int
	IsPrivate               *bool
	JoinByRequest           *bool
	IsJoined                *bool
	ParticipantCount        *int
	AreNotificationsEnabled *bool
	LastMessage             *message.Message
}
