package handler

import (
	"fmt"
	"homework_44/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
var Users []models.User
func GetUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for _, user := range Users{
		if user.Id == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
