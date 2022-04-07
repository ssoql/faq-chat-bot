package main

import (
	"github.com/ssoql/faq-chat-bot/src/api/app"
	"github.com/ssoql/faq-chat-bot/src/api/services"
)

func main() {
	go services.ChatService.Run()
	app.StartApp()
}
