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
	var users []handler.GetUserID.Users
	router := SetupRouter()

	users = append(users, models.User{Id: 1, Name: "Jane Doe", Email: "jane.doe@example.com"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetchedUser models.User
	json.Unmarshal(w.Body.Bytes(), &fetchedUser)
	assert.Equal(t, "Jane Doe", fetchedUser.Name)
	assert.Equal(t, "jane.doe@example.com", fetchedUser.Email)
	assert.Equal(t, 1, fetchedUser.Id)
}
