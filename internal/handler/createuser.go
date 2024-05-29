package handler

import (
	"fmt"
	"homework_44/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	var (
		users []models.User
		user models.User
		nextID=1
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	user.Id = nextID
	nextID++
	users = append(users, user)

	c.JSON(http.StatusOK, user)
}
