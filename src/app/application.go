package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-chat-bot/src/config"
	"github.com/ssoql/faq-chat-bot/src/datasources/faqs_db"
	"github.com/ssoql/faq-chat-bot/src/models/faqs"
	"github.com/ssoql/faq-chat-bot/src/services"
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
	services.FaqService.InitializeDemoFaqs()
	mapUrls()

	if err := router.Run(config.GetPort()); err != nil {
		panic(err)
	}
}
