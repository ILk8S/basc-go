package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", func(context *gin.Context) {

	})
	server.Run("localhost:3000")
}
