package chatsy

type ChatResponse struct {
	ID                      *int                `json:"id"`
	Type                    *string             `json:"type"`
	Title                   *string             `json:"title"`
	Description             *string             `json:"description"`
	InviteHash              *string             `json:"inviteHash"`
	JoinByLinkCount         *int                `json:"joinByLinkCount"`
	IsPrivate               *bool               `json:"isPrivate"`
	JoinByRequest           *bool               `json:"joinByRequest"`
	IsJoined                *bool               `json:"isJoined"`
	MemberCount             *int                `json:"memberCount"`
	AreNotificationsEnabled *bool               `json:"areNotificationsEnabled"`
	LastMessage             *GetMessageResponse `json:"lastMessage"`
}

type GetChatByIDRequest struct {
	ID *int `json:"id"`
}

type GetChatByIDResponse struct {
	*ChatResponse
}

type GetUserChatsPageRequest struct {
	UserID     *int `json:"userId"`
	Pagination Pagination
}

type GetUserChatsPageResponse struct {
	*PageResponse[*ChatResponse]
}

type GetUserChatsRequest struct {
	UserID *int `json:"userId"`
}

type GetUserChatsResponse struct {
	*ChatResponse
}

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
	MemberCount             *int                `json:"memberCount"`
	AreNotificationsEnabled *bool               `json:"areNotificationsEnabled"`
	LastMessage             *GetMessageResponse `json:"lastMessage"`
}
