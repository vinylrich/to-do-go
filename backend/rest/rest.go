package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) {
	router := gin.Default()
	router.GET("/", GetMainPage)
	router.GET("/ping", GetPing)

	router.Run(":8000")
}
