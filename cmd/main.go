package main

import (
	"github.com/gin-gonic/gin"
	"homework_44/internal/handler"
)

func main() {
	router := gin.Default()
	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUserID)
	router.GET("/createusers", handler.CreateUsers)

	router.Run(":8080")
}
