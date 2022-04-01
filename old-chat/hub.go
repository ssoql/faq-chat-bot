package main

import "log"

type Hub struct {
	chatbot Assistant
	clients map[*Client]bool

	broadcastmsg chan *ClientMessage
	register     chan *Client
	unregister   chan *Client
}

func startChatHub(hub *Hub) {

	go hub.Run()
}

func (h *Hub) ChatBot() Assistant {
	return h.chatbot
}

func (h *Hub) SendMessage(client *Client, message []byte) {

	client.send <- message
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			log.Println("new client redistration")
			h.clients[client] = true
			greeting := h.chatbot.Greeting()
			h.SendMessage(client, []byte(greeting))

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case clientmsg := <-h.broadcastmsg:
			client := clientmsg.client
			question := "You: " + string(clientmsg.message)
			h.SendMessage(client, []byte(question))
			reply := h.chatbot.Reply(string(clientmsg.message))
			h.SendMessage(client, []byte(reply))
		}
	}
}

func NewHub(chatbot Assistant) *Hub {

	return &Hub{
		chatbot:      chatbot,
		broadcastmsg: make(chan *ClientMessage),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
	}
}
