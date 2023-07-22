package main

import (
	"github.com/gorilla/websocket"
)

// client represents a chat user
type client struct {
	// socket represents websocket for this client
	socket *websocket.Conn

	// channel to receive messages from other client
	receive chan []byte

	// roon in which this client is chatting in
	room *room
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		// send the message to the channel
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
