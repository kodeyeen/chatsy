package websocket

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	conn     *websocket.Conn
	connMngr *ConnManager

	chatroom string
	// egress is used to avoid concurrent writes on the websocket connection
	Egress chan Event
}

func NewClient(conn *websocket.Conn, connMngr *ConnManager) *Client {
	return &Client{
		conn:     conn,
		connMngr: connMngr,
		Egress:   make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		// cleanup connection
		c.connMngr.removeClient(c)
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
			continue
		}

		err = c.connMngr.routeEvent(request, c)
		if err != nil {
			log.Println("error handeling message: ", err)
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.connMngr.removeClient(c)
	}()

	ticker := time.NewTicker(pingInterval)

	for {
		select {
		case message, ok := <-c.Egress:
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

func (c *Client) pongHandler(pongMsg string) error {
	log.Println("pong", pongMsg)
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}
