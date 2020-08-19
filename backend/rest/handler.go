package rest

import (
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	GetMainPage(c *gin.Context)
	GetPing(c *gin.Context)
}
type Handler struct {
	h HandlerInterface
}

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func GetMainPage(c *gin.Context) {
	c.JSON(200, gin.H{"Main": "Page"})
}
