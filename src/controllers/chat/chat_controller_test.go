package chat

import (
	"github.com/ssoql/faq-chat-bot/src/mocks"
	"github.com/ssoql/faq-chat-bot/src/services"
	"github.com/ssoql/faq-chat-bot/src/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunWebSocketNoError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/ws", nil)
	c, _ := test_utils.GetContextMock(request, response)

	services.ChatService = &mocks.ChatServiceMock{}

	RunWebSocket(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
}
