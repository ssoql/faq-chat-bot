package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-chat-bot/src/models/chats"
)

type ChatServiceMock struct {
}

func (s *ChatServiceMock) Run()                           {}
func (s *ChatServiceMock) Register(*chats.Client)         {}
func (s *ChatServiceMock) Unregister(*chats.Client)       {}
func (s *ChatServiceMock) Broadcast(*chats.ClientMessage) {}
func (s *ChatServiceMock) ServeWs(c *gin.Context)         {}
