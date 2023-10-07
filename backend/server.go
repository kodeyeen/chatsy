package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

type Server struct {
	sync.RWMutex

	events   chan Event
	clients  map[*Client]bool
	handlers map[string]EventHandler
}

func NewServer(ctx context.Context) *Server {
	s := &Server{
		clients:  make(map[*Client]bool),
		handlers: make(map[string]EventHandler),
	}

	s.setupEventHandlers()

	return s
}

func (s *Server) setupEventHandlers() {
	s.handlers[EventSendMessage] = SendMessage
	s.handlers[EventChangeRoom] = ChatRoomHandler
}

func (s *Server) routeEvent(event Event, c *Client) error {
	handler, ok := s.handlers[event.Type]

	if !ok {
		return errors.New("there is no such event type")
	}

	if err := handler(event, c); err != nil {
		return err
	}

	return nil
}

func SendMessage(event Event, c *Client) error {
	var chatevent SendMessageEvent

	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	var broadMessage NewMessageEvent

	broadMessage.SentAt = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.From = chatevent.From

	data, err := json.Marshal(broadMessage)

	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	outgoinEvent := Event{
		Payload: data,
		Type:    EventNewMessage,
	}

	for client := range c.server.clients {
		if client.chatroom == c.chatroom {
			client.egress <- outgoinEvent
		}
	}

	return nil
}

func ChatRoomHandler(event Event, c *Client) error {
	var changeRoomEvent ChangeRoomEvent

	if err := json.Unmarshal(event.Payload, &changeRoomEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	c.chatroom = changeRoomEvent.Name

	return nil
}

func (s *Server) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")

	// upgrade regular http connection into websocket
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, s)
	s.addClient(client)

	// start client processes
	go client.receiveMessages()
	go client.sendMessages()
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	type userLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req userLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "percy" && req.Password == "123" {
		w.WriteHeader(http.StatusOK)
		// w.Write(data)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func (s *Server) addClient(client *Client) {
	s.Lock()
	defer s.Unlock()

	s.clients[client] = true
}

func (s *Server) removeClient(client *Client) {
	s.Lock()
	defer s.Unlock()

	_, ok := s.clients[client]

	if ok {
		client.conn.Close()
		delete(s.clients, client)
	}
}

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	switch origin {
	case "http://localhost:8080":
		return true
	default:
		return false
	}
}
