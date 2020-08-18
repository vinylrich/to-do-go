package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8000")
}
