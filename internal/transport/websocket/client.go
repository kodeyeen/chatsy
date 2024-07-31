package websocket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type client struct {
	conn *websocket.Conn
	mgr  *Manager

	chatroom string
	// egress is used to avoid concurrent writes on the websocket connection
	egress chan Event
	usrID  int
}

func newClient(conn *websocket.Conn, mgr *Manager, usrID int) *client {
	return &client{
		conn:   conn,
		mgr:    mgr,
		egress: make(chan Event),
		usrID:  usrID,
	}
}

func (c *client) readMessages() {
	// cleanup connection
	defer func() {
		c.mgr.removeClient(c, userGroupName(c.usrID))
	}()

	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		log.Println(err)
		return
	}

	c.conn.SetReadLimit(512)
	c.conn.SetPongHandler(c.pongHandler)

	for {
		var request Event

		err := c.conn.ReadJSON(&request)
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}

		if err := c.mgr.routeEvent(request, c); err != nil {
			log.Println("error handling message: ", err)
		}
	}
}

func (c *client) writeMessages() {
	defer func() {
		c.mgr.removeClient(c, userGroupName(c.usrID))
	}()

	ticker := time.NewTicker(pingInterval)

	for {
		select {
		case message, ok := <-c.egress:
			// if channel is closed
			if !ok {
				// send message to the client, that we're closing the connection
				err := c.conn.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					// if sending is failed, the connection has probably been closed
					log.Println("connection closed: ")
				}

				return
			}

			err := c.conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("message sent")
		case <-ticker.C:
			log.Println("ping")

			err := c.conn.WriteMessage(websocket.PingMessage, []byte(""))
			if err != nil {
				log.Println("writemsg err", err)
				return
			}
		}
	}
}

func (c *client) pongHandler(pongMsg string) error {
	log.Println("pong", pongMsg)
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}
