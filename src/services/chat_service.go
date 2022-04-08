package services

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-chat-bot/src/models/chats"

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
	clients      map[*chats.Client]bool
	broadcastmsg chan *chats.ClientMessage
	register     chan *chats.Client
	unregister   chan *chats.Client
	Testing      chan string
}

type ChatServiceInterface interface {
	Run()
	Register(*chats.Client)
	Unregister(*chats.Client)
	Broadcast(*chats.ClientMessage)
}

func init() {
	ChatService = &chatService{
		broadcastmsg: make(chan *chats.ClientMessage),
		register:     make(chan *chats.Client),
		unregister:   make(chan *chats.Client),
		clients:      make(map[*chats.Client]bool),
	}
}

func (ch *chatService) Unregister(c *chats.Client) {
	ch.unregister <- c
}

func (ch *chatService) Broadcast(msg *chats.ClientMessage) {
	ch.broadcastmsg <- msg
}

func (ch *chatService) Register(c *chats.Client) {
	ch.register <- c
	log.Printf("registered in hub \n")
}

func (ch *chatService) RegisterRead(c *chats.Client) {
	ch.register <- c
}

func (ch *chatService) Run() {
	log.Println("start listening")
	defer log.Println("end listening")

	for {
		select {
		case client := <-ch.register:
			ch.clients[client] = true
			greeting := BotService.Greeting()
			client.SendMsg([]byte(greeting))

		case client := <-ch.unregister:
			if _, ok := ch.clients[client]; ok {
				delete(ch.clients, client)
				client.CloseSend()
			}
		case clientmsg := <-ch.broadcastmsg:
			client := clientmsg.GetClient()

			reply := BotService.Reply(string(clientmsg.GetMessage()))
			client.SendMsg([]byte(reply))
		}
	}
}

func ServeWs(c *gin.Context) {
	println(ChatService)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := chats.NewClient(conn)
	ChatService.Register(client)

	go writePump(client)
	readPump(client)
}
