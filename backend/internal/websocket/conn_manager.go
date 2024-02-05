package websocket

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/kodeyeen/chatsy/internal/ticket"
	"github.com/kodeyeen/chatsy/internal/user"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	switch origin {
	case "http://127.0.0.1:5173", "http://localhost:5173":
		return true
	default:
		return false
	}
}

type ticketFinder interface {
	FindByID(ctx context.Context, id string) (*ticket.Ticket, error)
}

type userRepository interface {
	Add(ctx context.Context, u *user.User) (int, error)
	FindByID(ctx context.Context, id int) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}

type ConnManager struct {
	sync.RWMutex

	clients      map[*Client]struct{}
	handlers     map[EventType]EventHandler
	ticketFinder ticketFinder
	userRepo     userRepository
}

func NewConnManager(ticketFinder ticketFinder, userRepo userRepository) *ConnManager {
	m := &ConnManager{
		clients:      make(map[*Client]struct{}),
		handlers:     make(map[EventType]EventHandler),
		ticketFinder: ticketFinder,
		userRepo:     userRepo,
	}

	m.setupEventHandlers()

	return m
}

func (m *ConnManager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessage
	m.handlers[EventChangeRoom] = ChatRoomHandler
}

func (m *ConnManager) Handle(evtType EventType, h EventHandler) {
	m.handlers[evtType] = h
}

func (m *ConnManager) ServeWS(w http.ResponseWriter, r *http.Request) {
	ticketID := r.URL.Query().Get("ticketID")
	if ticketID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	t, err := m.ticketFinder.FindByID(ctx, ticketID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("new connection")

	// upgrade regular http connection into websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = m.userRepo.FindByID(ctx, t.UserID)

	client := NewClient(conn, m)
	m.addClient(client)

	// start client processes
	go client.readMessages()
	go client.writeMessages()
}

func (m *ConnManager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = struct{}{}
}

func (m *ConnManager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	_, ok := m.clients[client]

	if ok {
		client.conn.Close()
		delete(m.clients, client)
	}
}

func (m *ConnManager) routeEvent(event Event, c *Client) error {
	handler, ok := m.handlers[event.Type]
	if !ok {
		return errors.New("there is no such event type")
	}

	if err := handler(event, c); err != nil {
		return err
	}

	return nil
}
