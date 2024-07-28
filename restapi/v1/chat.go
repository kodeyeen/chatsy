package restapi

type GetChatResponse struct {
	ID                      *int                `json:"id"`
	Type                    *string             `json:"type"`
	Title                   *string             `json:"title"`
	Description             *string             `json:"description"`
	InviteHash              *string             `json:"inviteHash"`
	JoinByLinkCount         *int                `json:"joinByLinkCount"`
	IsPrivate               *bool               `json:"isPrivate"`
	JoinByRequest           *bool               `json:"joinByRequest"`
	IsJoined                *bool               `json:"isJoined"`
	ParticipantCount        *int                `json:"participantCount"`
	AreNotificationsEnabled *bool               `json:"areNotificationsEnabled"`
	LastMessage             *GetMessageResponse `json:"lastMessage"`
}
