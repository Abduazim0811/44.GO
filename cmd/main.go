package main

import (
	"github.com/gin-gonic/gin"
	"homework_44/internal/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/users", handler.GetUser)
	router.GET("/users/:id", handler.GetUserID)
	router.POST("/createusers", handler.CreateUser)

	router.Run(":8080")
}
