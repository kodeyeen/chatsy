package websocket

import (
	"encoding/json"
	"time"

	"github.com/kodeyeen/chatsy/v1"
)

type EventType string

const (
	EventConnected       EventType = "connected"
	EventDisconnected    EventType = "disconnected"
	EventOpenChat        EventType = "open_chat"
	EventChatOpened      EventType = "chat_opened"
	EventSendMessages    EventType = "send_messages"
	EventNewMessages     EventType = "new_messages"
	EventFetchMessages   EventType = "fetch_messages"
	EventMessagesFetched EventType = "messages_fetched"

	EventSendMessage EventType = "send_message"
	EventNewMessage  EventType = "new_message"
	EventChangeRoom  EventType = "change_room"
)

type Event struct {
	Type    EventType       `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type HandlerFunc func(event Event, cl *client) error

type ConnectedEvent struct {
	Chats chatsy.GetUserChatsPageResponse `json:"chats"`
}

type OpenChatEvent struct {
	ChatID int `json:"chatId"`
}

type ChatOpenedEvent struct {
	Messages chatsy.PageResponse[*chatsy.GetMessageResponse] `json:"messages"`
}

type SendMessagesEvent struct {
	ChatID   int                            `json:"chatId"`
	Messages []*chatsy.CreateMessageRequest `json:"messages"`
}

type NewMessagesEvent struct {
	ChatID   int                          `json:"chatId"`
	Messages []*chatsy.GetMessageResponse `json:"messages"`
}

type FetchMessagesEvent struct {
	ChatID int `json:"chatId"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type LoadMessagesEvent struct {
	ChatID   int                                             `json:"chatId"`
	Messages chatsy.PageResponse[*chatsy.GetMessageResponse] `json:"messages"`
}

//

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type NewMessageEvent struct {
	SendMessageEvent
	SentAt time.Time `json:"sentAt"`
}

type ChangeRoomEvent struct {
	Name string `json:"name"`
}

func SendMessage(event Event, cl *client) error {
	// var chatevent SendMessageEvent
	// err := json.Unmarshal(event.Payload, &chatevent)
	// if err != nil {
	// 	return
	// }

	// var broadMessage NewMessageEvent
	// broadMessage.Message = chatevent.Message
	// broadMessage.From = chatevent.From
	// broadMessage.SentAt = time.Now()

	// data, err := json.Marshal(broadMessage)
	// if err != nil {
	// 	return
	// }

	// outgoinEvent := Event{
	// 	Type:    EventNewMessage,
	// 	Payload: data,
	// }

	// for client := range m.clients {
	// 	if client.chatroom == c.chatroom {
	// 		client.Egress <- outgoinEvent
	// 	}
	// }

	return nil
}

func ChatRoomHandler(event Event, cl *client) error {
	var changeRoomEvent ChangeRoomEvent
	err := json.Unmarshal(event.Payload, &changeRoomEvent)
	if err != nil {
		return err
	}

	cl.chatroom = changeRoomEvent.Name

	return nil
}
