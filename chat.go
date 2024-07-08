package chatsy

type ChatType string

const (
	ChatTypeDialog ChatType = "dialog"
	ChatTypeGroup  ChatType = "group"
)

type Chat struct {
	ID                      *int
	Type                    *ChatType
	Title                   *string
	Description             *string
	InviteHash              *string
	JoinByLinkCount         *int
	IsPrivate               *bool
	JoinByRequest           *bool
	IsJoined                *bool
	ParticipantCount        *int
	AreNotificationsEnabled *bool
	LastMessage             *Message
}
