package handler

import (
	"fmt"
	"homework_44/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	users  []models.User
	nextID = 1
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for _, user := range users {
		if user.Id == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	user.Id = nextID
	nextID++
	users = append(users, user)

	c.JSON(http.StatusOK, user)
}
