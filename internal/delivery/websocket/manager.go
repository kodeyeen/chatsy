package websocket

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)

type Manager struct {
	sync.RWMutex

	groups   map[string]map[*client]struct{}
	handlers map[EventType]HandlerFunc
}

func NewManager(handler *EventHandler) *Manager {
	m := &Manager{
		groups:   make(map[string]map[*client]struct{}),
		handlers: make(map[EventType]HandlerFunc),
	}

	m.setupEventHandlers()

	return m
}

func (m *Manager) setupEventHandlers() {
	// m.handlers[EventOpenChat] = m.handler.onOpenChat
	// m.handlers[EventSendMessages] = m.handler.onSendMessages
	// m.handlers[EventFetchMessages] = m.handler.onFetchMessages

	// m.handlers[EventSendMessage] = SendMessage
	// m.handlers[EventChangeRoom] = ChatRoomHandler
}

func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// upgrade regular http connection into websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("new connection")

	cl := newClient(conn, m, userID)
	m.addClient(cl, userGroupName(userID))

	// start client processes
	go cl.readMessages()
	go cl.writeMessages()
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

func (m *Manager) routeEvent(evt Event, cl *client) error {
	handler, ok := m.handlers[evt.Type]
	if !ok {
		// not found handler
		log.Println("not found")
		return errors.New("there is no such event type")
	}

	return handler(evt, cl)
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
