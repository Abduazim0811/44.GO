package testing

import (
	"homework_44/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUserID)
	router.POST("/users", handler.CreateUsers)
	return router
}
