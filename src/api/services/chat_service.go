package services

import (
	"fmt"
	"github.com/ssoql/faq-chat-bot/src/api/models"

	"log"
)

var (
	//clients      map[*models.Client]bool
	//broadcastmsg chan *models.ClientMessage
	//register     chan *models.Client
	//unregister   chan *models.Client
	ChatService ChatServiceInterface
)

type chatService struct {
	clients      map[*models.Client]bool
	broadcastmsg chan *models.ClientMessage
	register     chan *models.Client
	unregister   chan *models.Client
	Testing      chan string
}

type ChatServiceInterface interface {
	Run()
	Register(*models.Client)
	Unregister(*models.Client)
	Broadcast(*models.ClientMessage)
}

func init() {
	ChatService = &chatService{
		broadcastmsg: make(chan *models.ClientMessage),
		register:     make(chan *models.Client),
		unregister:   make(chan *models.Client),
		clients:      make(map[*models.Client]bool),
	}
}

func (ch *chatService) Unregister(c *models.Client) {
	ch.unregister <- c
}

func (ch *chatService) Broadcast(msg *models.ClientMessage) {
	ch.broadcastmsg <- msg
}

func (ch *chatService) Register(c *models.Client) {
	ch.register <- c
	log.Printf("registered in hub \n")
}

func (ch *chatService) RegisterRead(c *models.Client) {
	ch.register <- c
}

func (ch *chatService) Run() {
	log.Println("start listening")
	defer log.Println("end listening")

	for {
		select {
		case client := <-ch.register:
			log.Println("new client registration")
			fmt.Printf("new CLIENT: %v", client)
			ch.clients[client] = true
			log.Printf("Users on chat: %v", len(ch.clients))
			greeting := BotService.Greeting()
			client.SendMsg([]byte(greeting))

		case client := <-ch.unregister:
			log.Println("client unregister")
			if _, ok := ch.clients[client]; ok {
				delete(ch.clients, client)
				client.CloseSend()
			}
		case clientmsg := <-ch.broadcastmsg:
			log.Println("new broadcast")
			client := clientmsg.GetClient()
			question := "You: " + string(clientmsg.GetMessage())
			client.SendMsg([]byte(question))

			reply := BotService.Reply(string(clientmsg.GetMessage()))
			client.SendMsg([]byte(reply))
		}
	}
}
