package status

import (
	"github.com/ssoql/faq-chat-bot/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "OK", statusString)
}

func TestCheck(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/status", nil)
	c := test_utils.GetContextMock(request, response)

	Check(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "OK", response.Body.String())
}
