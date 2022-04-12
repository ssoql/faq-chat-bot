package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-chat-bot/src/services"
)

func RunWebSocket(c *gin.Context) {
	services.ChatService.ServeWs(c)
}
