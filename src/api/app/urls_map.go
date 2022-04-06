package app

import (
	"github.com/ssoql/faq-chat-bot/src/api/controllers/chat"
	"github.com/ssoql/faq-chat-bot/src/api/controllers/faq"
	"github.com/ssoql/faq-chat-bot/src/api/controllers/home"
	"github.com/ssoql/faq-chat-bot/src/api/controllers/status"
)

func mapUrls() {
	router.GET("/", home.ShowHomePage)
	router.GET("/ws", chat.RunWebSocket)
	router.GET("/status", status.Check)
	router.POST("/faq", faq.Create)
	router.PATCH("/faq/:faq_id", faq.Update)
	//router.POST("/repository", repositories.CreateRepo)
	//router.POST("/repositories", repositories.CreateRepos)
}
