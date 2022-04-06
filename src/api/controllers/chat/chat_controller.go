package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-chat-bot/src/api/services"
)

func RunWebSocket(c *gin.Context) {
	services.ServeWs(c)
}
