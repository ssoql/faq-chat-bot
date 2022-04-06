package main

import (
	"github.com/ssoql/faq-chat-bot/src/api/app"
	"github.com/ssoql/faq-chat-bot/src/api/config"
	"github.com/ssoql/faq-chat-bot/src/api/services"
	"time"
)

func main() {
	if !config.IsProduction() {
		println("Sleep start")
		time.Sleep(10000)
		println("Sleep end")
	}
	go services.ChatService.Run()
	app.StartApp()
}
