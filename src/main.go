package main

import (
	"github.com/ssoql/faq-chat-bot/src/app"
	"github.com/ssoql/faq-chat-bot/src/services"
)

func main() {
	// initialize chat goroutine
	go services.ChatService.Run()
	// start web app
	app.StartApp()
}
