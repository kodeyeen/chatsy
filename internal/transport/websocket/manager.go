package websocket

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	sync.RWMutex

	upgrader *websocket.Upgrader
	groups   map[string]map[*client]struct{}
	handlers map[EventType]HandlerFunc

	handler *EventHandler
	ctx     context.Context
}

func NewManager(handler *EventHandler) *Manager {
	m := &Manager{
		groups:   make(map[string]map[*client]struct{}),
		handlers: make(map[EventType]HandlerFunc),

		handler: handler,
		ctx:     context.Background(),
	}

	m.upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     m.checkOrigin,
	}

	m.setupEventHandlers()

	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventOpenChat] = m.handler.onOpenChat
	m.handlers[EventSendMessages] = m.handler.onSendMessages
	m.handlers[EventFetchMessages] = m.handler.onFetchMessages

	m.handlers[EventSendMessage] = SendMessage
	m.handlers[EventChangeRoom] = ChatRoomHandler
}

func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// upgrade regular http connection into websocket
	conn, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("new connection")

	cl := newClient(conn, userID)

	// start client processes
	go cl.readMessages(m)
	go cl.writeMessages(m)

	m.connect(cl)
}

func (m *Manager) connect(cl *client) {
	connectedEvent := Event{
		Type: EventConnected,
	}

	m.handler.onConnect(connectedEvent, cl, m)
}

func (m *Manager) disconnect(cl *client) {
	disconnectedEvent := Event{
		Type: EventDisconnected,
	}

	m.handler.onDisconnect(disconnectedEvent, cl, m)
}

func (m *Manager) addClient(cl *client, groupName string) {
	m.Lock()
	defer m.Unlock()

	_, ok := m.groups[groupName]

	if !ok {
		m.groups[groupName] = make(map[*client]struct{})
	}

	m.groups[groupName][cl] = struct{}{}
}

func (m *Manager) removeClient(cl *client, groupName string) {
	m.Lock()
	defer m.Unlock()

	group, ok := m.groups[groupName]

	if !ok {
		return
	}

	_, ok = group[cl]

	if ok {
		cl.conn.Close()
		delete(group, cl)
	}
}

func (m *Manager) routeEvent(evt Event, cl *client) {
	handler, ok := m.handlers[evt.Type]
	if !ok {
		// not found handler
		log.Println("not found")
		return
	}

	handler(evt, cl, m)
}

func (m *Manager) sendEvent(evt Event, groupName string) {
	group := m.groups[groupName]

	for cl := range group {
		cl.sendEvent(evt)
	}
}

func (m *Manager) checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	switch origin {
	case "http://127.0.0.1:5173", "http://localhost:5173":
		return true
	default:
		return false
	}
}
