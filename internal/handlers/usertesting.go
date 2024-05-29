package handler

import (
	"bytes"
	"encoding/json"
	"homework_44/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", GetUser)
	router.GET("/users/:id", GetUserID)
	router.POST("/users", CreateUser)
	return router
}

func TestGetUsers(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	user := models.User{Name: "Abduazim", Email: "abduazim@gmail.com"}
	jsonValue, _ := json.Marshal(user)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var createdUser models.User
	json.Unmarshal(w.Body.Bytes(), &createdUser)
	assert.Equal(t, "Abduazim", createdUser.Name)
	assert.Equal(t, "abduazim@gmail.com", createdUser.Email)
	assert.Equal(t, 1, createdUser.Id)
}

func TestGetUserByID(t *testing.T) {
	router := setupRouter()

	users = append(users, models.User{Id: 1, Name: "Abduazim", Email: "abduazim@gmail.com"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetchedUser models.User
	json.Unmarshal(w.Body.Bytes(), &fetchedUser)
	assert.Equal(t, "Abduazim", fetchedUser.Name)
	assert.Equal(t, "abduazim@gmail.com", fetchedUser.Email)
	assert.Equal(t, 1, fetchedUser.Id)
}

func TestGetUserByID_NotFound(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
