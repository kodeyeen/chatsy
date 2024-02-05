package chat

type Type string

const (
	TypePersonal Type = "personal"
	TypeGroup    Type = "group"
)

type Chat struct {
	ID              int
	Type            Type
	Title           string
	Description     string
	InviteHash      string
	JoinByLinkCount int
	IsPrivate       bool
	JoinByRequest   bool
}
