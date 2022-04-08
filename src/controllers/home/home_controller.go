package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Simple Chat Bot",
	})
}
