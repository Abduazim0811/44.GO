package handler

import (
	"homework_44/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	c.JSON(http.StatusOK, users)
}
