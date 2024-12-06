package chatsy

type Client struct {
	messages *MessageService
}

func NewClient() *Client {
	return &Client{
		messages: &MessageService{},
	}
}
