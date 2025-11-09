package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ILk8S/basc-go/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	hdl := web.NewUserHandler()
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "/users")
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		fmt.Println("这是Middleware")
	})
	hdl.RegisterRoutes(server)
	server.Run(":8899")
}
