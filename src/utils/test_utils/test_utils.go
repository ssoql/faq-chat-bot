package test_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func GetContextMock(request *http.Request, response *httptest.ResponseRecorder) (*gin.Context, *gin.Engine) {
	c, r := gin.CreateTestContext(response)
	c.Request = request
	return c, r
}
