package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kodeyeen/chatsy/internal/websocket"
)

type wsHandler struct {
	svc service
}

func NewWsHandler(svc service) *wsHandler {
	return &wsHandler{
		svc: svc,
	}
}

func (h *wsHandler) FetchChats(event websocket.Event, c *websocket.Client) error {
	var payload FetchChatsEvent
	err := json.Unmarshal(event.Payload, &payload)
	if err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	ctx := context.Background()
	page, err := h.svc.GetUserChats(ctx, 1, payload.Limit, payload.Offset)

	out := ChatsFetchedEvent{
		Chats: page,
	}

	data, err := json.Marshal(out)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	outgoingEvent := websocket.Event{
		Type:    websocket.EventChatsFetched,
		Payload: data,
	}

	c.Egress <- outgoingEvent

	return nil
}
