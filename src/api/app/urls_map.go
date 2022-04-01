package app

import (
	"github.com/ssoql/faq-chat-bot/src/api/controllers/home"
	"github.com/ssoql/faq-chat-bot/src/api/controllers/status"
	"github.com/ssoql/faq-chat-bot/src/api/services"
)

func mapUrls() {
	router.GET("/", home.ShowHomePage)
	router.GET("/ws", services.ServeWs)
	router.GET("/status", status.Check)
	//router.POST("/repository", repositories.CreateRepo)
	//router.POST("/repositories", repositories.CreateRepos)
}
