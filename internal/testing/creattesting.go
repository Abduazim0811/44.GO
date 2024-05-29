package testing

import (
	"bytes"
	"encoding/json"
	"homework_44/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/go-playground/assert/v2"
)

func TestCreateUser(t *testing.T) {
	router := SetupRouter()
	user := models.User{Name: "John Doe", Email: "john.doe@example.com"}
	jsonValue, _ := json.Marshal(user)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var createdUser models.User
	json.Unmarshal(w.Body.Bytes(), &createdUser)
	assert.Equal(t, "John Doe", createdUser.Name)
	assert.Equal(t, "john.doe@example.com", createdUser.Email)
	assert.Equal(t, 1, createdUser.Id)
}
