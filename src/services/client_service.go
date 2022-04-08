package services

import (
	"bytes"
	"github.com/gorilla/websocket"
	"github.com/ssoql/faq-chat-bot/src/models/chats"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline  = []byte{'\n'}
	space    = []byte{' '}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func readPump(c *chats.Client) {
	log.Println("start read pump")

	defer func() {
		ChatService.Unregister(c)
		c.CloseConnection()
	}()
	c.Connection().SetReadLimit(maxMessageSize)
	c.Connection().SetReadDeadline(time.Now().Add(pongWait))
	c.Connection().SetPongHandler(func(string) error {
		c.Connection().SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.Connection().ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error:%v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//c.hub.broadcast <- message
		clientmsg := chats.NewClientMessage(c, message)
		ChatService.Broadcast(clientmsg)
	}
}

func writePump(c *chats.Client) {
	log.Println("start write pump")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.CloseConnection()
	}()
	for {
		select {
		case message, ok := <-c.Send():
			//c.Connection().SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Connection().SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Println("Error setting Deadline to connection")
			}
			if !ok {
				// The chat closed the channel.
				if err := c.Connection().WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Println("Error writing the message and closing the writer")
				}
				return
			}

			w, err := c.Connection().NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			if _, err := w.Write(message); err != nil {
				log.Println("Error writing to websocket")
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.Send())
			for i := 0; i < n; i++ {
				if _, err := w.Write(newline); err != nil {
					log.Println("Error writing a new line to websocket")
				}
				if _, err := w.Write(<-c.Send()); err != nil {
					log.Println("Error writing data to websocket")
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:

			if err := c.Connection().SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Println("Error setting Deadline to connection")
			}
			if err := c.Connection().WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
