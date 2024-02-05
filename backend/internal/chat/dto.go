package chat

import "github.com/kodeyeen/chatsy/internal/api"

type GetDTO struct {
	ID              int    `json:"id" db:"id"`
	Type            Type   `json:"type" db:"type"`
	Title           string `json:"title" db:"title"`
	Description     string `json:"description" db:"description"`
	InviteHash      string `json:"inviteHash" db:"invite_hash"`
	JoinByLinkCount int    `json:"joinByLinkCount" db:"join_by_link_count"`
	IsPrivate       bool   `json:"isPrivate" db:"is_private"`
	JoinByRequest   bool   `json:"joinByRequest" db:"join_by_request"`
}

type ChatPage struct {
	Items  []*GetDTO `json:"items"`
	Count  int       `json:"count"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}

type FetchChatsEvent struct {
	api.Pagination
}

type ChatsFetchedEvent struct {
	Chats *ChatPage `json:"chats"`
}
