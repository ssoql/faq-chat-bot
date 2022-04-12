package status

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const StatusString = "OK"

func Check(c *gin.Context) {
	c.String(http.StatusOK, StatusString)
}
