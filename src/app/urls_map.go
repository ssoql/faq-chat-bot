package app

import (
	"github.com/ssoql/faq-chat-bot/src/controllers/chat"
	"github.com/ssoql/faq-chat-bot/src/controllers/faq"
	"github.com/ssoql/faq-chat-bot/src/controllers/home"
	"github.com/ssoql/faq-chat-bot/src/controllers/status"
)

func mapUrls() {
	router.GET("/", home.ShowHomePage)
	router.GET("/ws", chat.RunWebSocket)
	router.GET("/status", status.Check)
	router.POST("/faq", faq.Create)
	router.GET("/faq/:faq_id", faq.Get)
	router.PATCH("/faq/:faq_id", faq.Update)
	router.DELETE("/faq/:faq_id", faq.Delete)
}
