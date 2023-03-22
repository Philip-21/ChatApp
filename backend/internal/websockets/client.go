package websockets

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// client establishes a connection
type Client struct {
	ID   string          //uniquely identifiably string for a particular connection
	Conn *websocket.Conn //a pointer to a websocket.Conn object
	Pool *Pool           //a pointer to the Pool which this client will be part of
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

//constantly listens for new messages coming through on
// this Client’s websocket connection.

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
