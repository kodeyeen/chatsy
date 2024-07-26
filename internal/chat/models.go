package chat

import (
	"github.com/kodeyeen/chatsy"
	"github.com/kodeyeen/chatsy/internal/message"
)

type GetResponse struct {
	ID                      *int                 `json:"id"`
	Type                    *chatsy.ChatType     `json:"type"`
	Title                   *string              `json:"title"`
	Description             *string              `json:"description"`
	InviteHash              *string              `json:"inviteHash"`
	JoinByLinkCount         *int                 `json:"joinByLinkCount"`
	IsPrivate               *bool                `json:"isPrivate"`
	JoinByRequest           *bool                `json:"joinByRequest"`
	IsJoined                *bool                `json:"isJoined"`
	ParticipantCount        *int                 `json:"participantCount"`
	AreNotificationsEnabled *bool                `json:"areNotificationsEnabled"`
	LastMessage             *message.GetResponse `json:"lastMessage"`
}
