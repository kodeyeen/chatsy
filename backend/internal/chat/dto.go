package chat

import "github.com/kodeyeen/chatsy/internal/message"

type GetDTO struct {
	ID                      *int            `json:"id"`
	Type                    *Type           `json:"type"`
	Title                   *string         `json:"title"`
	Description             *string         `json:"description"`
	InviteHash              *string         `json:"inviteHash"`
	JoinByLinkCount         *int            `json:"joinByLinkCount"`
	IsPrivate               *bool           `json:"isPrivate"`
	JoinByRequest           *bool           `json:"joinByRequest"`
	IsJoined                *bool           `json:"isJoined"`
	ParticipantCount        *int            `json:"participantCount"`
	AreNotificationsEnabled *bool           `json:"areNotificationsEnabled"`
	LastMessage             *message.GetDTO `json:"lastMessage"`
}
