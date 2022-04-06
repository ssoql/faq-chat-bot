package chats

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type ClientMessage struct {
	client  *Client
	message []byte
}

func (m *ClientMessage) GetClient() *Client {
	return m.client
}

func (m *ClientMessage) GetMessage() []byte {
	return m.message
}

func (m ClientMessage) SetClient(client *Client) {
	m.client = client
}

func (m ClientMessage) SetMessage(message []byte) {
	m.message = message
}

func (c *Client) SendMsg(msg []byte) {
	c.send <- msg
}

func (c *Client) Send() chan []byte {
	return c.send
}

func (c *Client) Connection() *websocket.Conn {
	return c.conn
}

func (c *Client) CloseSend() {
	close(c.send)
}

func (c *Client) CloseConnection() {
	if err := c.conn.Close(); err != nil {
		log.Println(err.Error())
	}
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		send: make(chan []byte, 256),
	}
}

func NewClientMessage(client *Client, message []byte) *ClientMessage {
	return &ClientMessage{
		client:  client,
		message: message,
	}
}
