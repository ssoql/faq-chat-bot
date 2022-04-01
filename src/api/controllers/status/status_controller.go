package status

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const statusString = "OK"

func Check(c *gin.Context) {
	c.String(http.StatusOK, statusString)
}
