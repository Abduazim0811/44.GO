package testing

import (
	"encoding/json"
	"homework_44/internal/models"
	"homework_44/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestingUserId(t *testing.T) {
	router := SetupRouter()

	handler.Users = append(handler.Users, models.User{Id: 1, Name: "Abduazim", Email: "abduazim@gmail.com"})

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
