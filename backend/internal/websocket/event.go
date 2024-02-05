package websocket

import (
	"encoding/json"
	"fmt"
	"time"
)

type EventType string

const (
	EventFetchChats   EventType = "fetch_chats"
	EventChatsFetched EventType = "chats_fetched"

	EventSendMessage EventType = "send_message"
	EventNewMessage  EventType = "new_message"
	EventChangeRoom  EventType = "change_room"
)

type Event struct {
	Type    EventType       `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

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

func SendMessage(event Event, c *Client) error {
	var chatevent SendMessageEvent
	err := json.Unmarshal(event.Payload, &chatevent)
	if err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	var broadMessage NewMessageEvent
	broadMessage.Message = chatevent.Message
	broadMessage.From = chatevent.From
	broadMessage.SentAt = time.Now()

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	outgoinEvent := Event{
		Type:    EventNewMessage,
		Payload: data,
	}

	for client := range c.connMngr.clients {
		if client.chatroom == c.chatroom {
			client.Egress <- outgoinEvent
		}
	}

	return nil
}

func ChatRoomHandler(event Event, c *Client) error {
	var changeRoomEvent ChangeRoomEvent
	err := json.Unmarshal(event.Payload, &changeRoomEvent)
	if err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	c.chatroom = changeRoomEvent.Name

	return nil
}
