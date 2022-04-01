package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-chat-bot/src/api/config"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
}

func StartApp() {
	mapUrls()
	if err := router.Run(config.GetPort()); err != nil {
		panic(err)
	}
}
