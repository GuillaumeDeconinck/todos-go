package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GuillaumeDeconinck/todos-go/internal/api"
	"github.com/GuillaumeDeconinck/todos-go/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandleListTodos(t *testing.T) {
	// Setup
	router := api.SetupApi()

	var uuidToDelete *string

	t.Run("List todos - should be empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todos", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// var received []models.Todo
		// json.Unmarshal(w.Body.Bytes(), &received)

		// var expected []models.Todo
		// assert.ElementsMatch(t, expected, received)
	})

	t.Run("Create a new todo", func(t *testing.T) {
		uuid, title, ownerUuid, state := uuid.New().String(), "Do the chores", uuid.New().String(), "ACTIVE"
		todo := models.Todo{Uuid: &uuid, Title: &title, OwnerUuid: &ownerUuid, State: &state}

		w := httptest.NewRecorder()
		json_data, _ := json.Marshal(todo)
		req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(json_data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
	})

	t.Run("List todos - should have one", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todos", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var received []models.Todo
		json.Unmarshal(w.Body.Bytes(), &received)

		uuidToDelete = received[0].Uuid

		// var expected []models.Todo
		// assert.ElementsMatch(t, expected, received)
	})

	t.Run("Delete todo", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/todos/"+*uuidToDelete, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 204, w.Code)
	})
}
