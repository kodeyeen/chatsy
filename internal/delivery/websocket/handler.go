package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kodeyeen/chatsy/v1"
)

type chatService interface {
	GetByID(ctx context.Context, req *chatsy.GetChatByIDRequest) (*chatsy.GetChatByIDResponse, error)
	GetAllForUser(ctx context.Context, req *chatsy.GetUserChatsRequest) ([]*chatsy.GetUserChatsResponse, error)
	GetByUserID(ctx context.Context, req *chatsy.GetUserChatsPageRequest) (*chatsy.GetUserChatsPageResponse, error)
}

type messageService interface {
	Create(ctx context.Context, req *chatsy.CreateMessageRequest, senderID int) (*chatsy.GetMessageResponse, error)
	GetByChatID(ctx context.Context, chatID int, limit, offset int) (*chatsy.PageResponse[*chatsy.GetMessageResponse], error)
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

func (h *EventHandler) onConnect(evt Event, cl *client) {
	ctx := context.Background()

	userChats, err := h.chatSvc.GetAllForUser(ctx, &chatsy.GetUserChatsRequest{
		UserID: &cl.usrID,
	})
	if err != nil {
		log.Println(err)
		return
	}

	cl.mgr.addClient(cl, userGroupName(cl.usrID))

	for _, userChat := range userChats {
		cl.mgr.addClient(cl, chatGroupName(*userChat.ID))
	}

	page, err := h.chatSvc.GetByUserID(ctx, &chatsy.GetUserChatsPageRequest{
		UserID: &cl.usrID,
		Pagination: chatsy.Pagination{
			Limit:  10,
			Offset: 0,
		},
	})
	if err != nil {
		log.Println("GetForUser err", err)
		return
	}

	_, err = json.Marshal(ConnectedEvent{
		Chats: *page,
	})
	if err != nil {
		return
	}

	// cl.sendEvent(Event{
	// 	Type:    EventConnected,
	// 	Payload: data,
	// })
}

func (h *EventHandler) onDisconnect(evt Event, cl *client) {
	ctx := context.Background()

	userChats, err := h.chatSvc.GetAllForUser(ctx, &chatsy.GetUserChatsRequest{
		UserID: &cl.usrID,
	})
	if err != nil {
		log.Println(err)
		return
	}

	cl.mgr.removeClient(cl, userGroupName(cl.usrID))

	for _, userChat := range userChats {
		cl.mgr.removeClient(cl, chatGroupName(*userChat.ID))
	}
}

func (h *EventHandler) onOpenChat(evt Event, cl *client) error {
	var openChatEvt OpenChatEvent
	err := json.Unmarshal(evt.Payload, &openChatEvt)
	if err != nil {
		return err
	}

	ctx := context.Background()

	page, err := h.msgSvc.GetByChatID(ctx, openChatEvt.ChatID, 7, 0)
	if err != nil {
		log.Println("GetForChat err", err)
		return err
	}

	_, err = json.Marshal(ChatOpenedEvent{
		Messages: *page,
	})
	if err != nil {
		return err
	}

	// cl.sendEvent(Event{
	// 	Type:    EventChatOpened,
	// 	Payload: data,
	// })

	return nil
}

func (h *EventHandler) onSendMessages(evt Event, cl *client) error {
	var sendMsgsEvt SendMessagesEvent
	err := json.Unmarshal(evt.Payload, &sendMsgsEvt)
	if err != nil {
		return err
	}

	ctx := context.Background()

	var msgs []*chatsy.GetMessageResponse

	// TODO: bulk create
	for _, dto := range sendMsgsEvt.Messages {
		msg, err := h.msgSvc.Create(ctx, dto, cl.usrID)
		if err != nil {
			log.Println("h.msgSvc.Create err", err)
			return err
		}

		msgs = append(msgs, msg)
	}

	_, err = json.Marshal(NewMessagesEvent{
		ChatID:   sendMsgsEvt.ChatID,
		Messages: msgs,
	})
	if err != nil {
		return err
	}

	// cl.mgr.sendEvent(Event{
	// 	Type:    EventNewMessages,
	// 	Payload: data,
	// }, chatGroupName(sendMsgsEvt.ChatID))

	return nil
}

func (h *EventHandler) onFetchMessages(evt Event, cl *client) error {
	var fetchMsgsEvt FetchMessagesEvent
	err := json.Unmarshal(evt.Payload, &fetchMsgsEvt)
	if err != nil {
		return err
	}

	ctx := context.Background()

	page, err := h.msgSvc.GetByChatID(ctx, fetchMsgsEvt.ChatID, fetchMsgsEvt.Limit, fetchMsgsEvt.Offset)
	if err != nil {
		return err
	}

	_, err = json.Marshal(LoadMessagesEvent{
		ChatID:   fetchMsgsEvt.ChatID,
		Messages: *page,
	})
	if err != nil {
		return err
	}

	// cl.mgr.sendEvent(Event{
	// 	Type:    EventMessagesFetched,
	// 	Payload: data,
	// }, chatGroupName(fetchMsgsEvt.ChatID))

	return nil
}
