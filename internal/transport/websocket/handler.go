package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kodeyeen/chatsy/restapi"
)

type chatService interface {
	GetByID(ctx context.Context, id int) (*restapi.GetChatResponse, error)
	GetAllForUser(ctx context.Context, userID int) ([]*restapi.GetChatResponse, error)
	GetForUser(ctx context.Context, userID int, limit, offset int) (*restapi.PageResponse[*restapi.GetChatResponse], error)
}

type messageService interface {
	Create(ctx context.Context, req *restapi.CreateMessageRequest, senderID int) (*restapi.GetMessageResponse, error)
	GetForChat(ctx context.Context, chatID int, limit, offset int) (*restapi.PageResponse[*restapi.GetMessageResponse], error)
}

func userGroupName(userID int) string {
	return fmt.Sprintf("user-%d", userID)
}

func chatGroupName(chatID int) string {
	return fmt.Sprintf("chat-%d", chatID)
}

type EventHandler struct {
	chatSvc chatService
	msgSvc  messageService
}

func NewEventHandler(
	chatSvc chatService,
	msgSvc messageService,
) *EventHandler {
	return &EventHandler{
		chatSvc: chatSvc,
		msgSvc:  msgSvc,
	}
}

func (h *EventHandler) onConnect(evt Event, cl *client, mng *Manager) {
	userChats, err := h.chatSvc.GetAllForUser(mng.ctx, cl.usrID)
	if err != nil {
		log.Println(err)
		return
	}

	mng.addClient(cl, userGroupName(cl.usrID))

	for _, userChat := range userChats {
		mng.addClient(cl, chatGroupName(*userChat.ID))
	}

	page, err := h.chatSvc.GetForUser(mng.ctx, cl.usrID, 10, 0)
	if err != nil {
		log.Println("GetForUser err", err)
		return
	}

	data, err := json.Marshal(ConnectedEvent{
		Chats: *page,
	})
	if err != nil {
		return
	}

	cl.sendEvent(Event{
		Type:    EventConnected,
		Payload: data,
	})
}

func (h *EventHandler) onDisconnect(evt Event, cl *client, mng *Manager) {
	userChats, err := h.chatSvc.GetAllForUser(mng.ctx, cl.usrID)
	if err != nil {
		log.Println(err)
		return
	}

	mng.removeClient(cl, userGroupName(cl.usrID))

	for _, userChat := range userChats {
		mng.removeClient(cl, chatGroupName(*userChat.ID))
	}
}

func (h *EventHandler) onOpenChat(evt Event, cl *client, mng *Manager) {
	var openChatEvt OpenChatEvent
	err := json.Unmarshal(evt.Payload, &openChatEvt)
	if err != nil {
		return
	}

	page, err := h.msgSvc.GetForChat(mng.ctx, openChatEvt.ChatID, 7, 0)
	if err != nil {
		log.Println("GetForChat err", err)
		return
	}

	data, err := json.Marshal(ChatOpenedEvent{
		Messages: *page,
	})
	if err != nil {
		return
	}

	cl.sendEvent(Event{
		Type:    EventChatOpened,
		Payload: data,
	})
}

func (h *EventHandler) onSendMessages(evt Event, cl *client, mng *Manager) {
	var sendMsgsEvt SendMessagesEvent
	err := json.Unmarshal(evt.Payload, &sendMsgsEvt)
	if err != nil {
		return
	}

	var msgs []*restapi.GetMessageResponse

	// TODO: bulk create
	for _, dto := range sendMsgsEvt.Messages {
		msg, err := h.msgSvc.Create(mng.ctx, dto, cl.usrID)
		if err != nil {
			log.Println("h.msgSvc.Create err", err)
			return
		}

		msgs = append(msgs, msg)
	}

	data, err := json.Marshal(NewMessagesEvent{
		ChatID:   sendMsgsEvt.ChatID,
		Messages: msgs,
	})
	if err != nil {
		return
	}

	mng.sendEvent(Event{
		Type:    EventNewMessages,
		Payload: data,
	}, chatGroupName(sendMsgsEvt.ChatID))
}

func (h *EventHandler) onFetchMessages(evt Event, cl *client, mng *Manager) {
	var fetchMsgsEvt FetchMessagesEvent
	err := json.Unmarshal(evt.Payload, &fetchMsgsEvt)
	if err != nil {
		return
	}

	page, err := h.msgSvc.GetForChat(mng.ctx, fetchMsgsEvt.ChatID, fetchMsgsEvt.Limit, fetchMsgsEvt.Offset)
	if err != nil {
		return
	}

	data, err := json.Marshal(LoadMessagesEvent{
		ChatID:   fetchMsgsEvt.ChatID,
		Messages: *page,
	})
	if err != nil {
		return
	}

	mng.sendEvent(Event{
		Type:    EventMessagesFetched,
		Payload: data,
	}, chatGroupName(fetchMsgsEvt.ChatID))
}
