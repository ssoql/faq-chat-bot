package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-chat-bot/src/api/config"
	"github.com/ssoql/faq-chat-bot/src/api/datasources/faqs_db"
	"github.com/ssoql/faq-chat-bot/src/api/models/faqs"
	"os"
	"path/filepath"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.LoadHTMLGlob(filepath.Join(os.Getenv("TMPL_PATH"), "*.tmpl"))
}

func StartApp() {
	faqs_db.MigrateData(&faqs.Faq{})
	mapUrls()

	if err := router.Run(config.GetPort()); err != nil {
		panic(err)
	}
}
