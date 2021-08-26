package api

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	todoToCreate := models.Todo{Uuid: uuid.New().String(), Title: "Do the chores", OwnerUuid: uuid.New().String(), State: "ACTIVE"}

	updatedTitle := "Buy groceries"

	// Will be the same as the Uuid above ^
	var uuidToDelete string

	t.Run("List todos - should be empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todos", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var received []models.Todo
		json.Unmarshal(w.Body.Bytes(), &received)

		var expected []models.Todo
		assert.ElementsMatch(t, expected, received)
	})

	t.Run("Get todo - should return a 404", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todos/1345", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)

		var received struct {
			Error   string `json:"error"`
			Message string `json:"message"`
		}
		json.Unmarshal(w.Body.Bytes(), &received)

		assert.Equal(t, "Not found", received.Error)
	})

	t.Run("Create a new todo", func(t *testing.T) {

		w := httptest.NewRecorder()
		json_data, _ := json.Marshal(todoToCreate)
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

		assert.Equal(t, 1, len(received))

		uuidToDelete = received[0].Uuid
	})

	t.Run("Get todo", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todos/"+uuidToDelete, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var received models.Todo
		json.Unmarshal(w.Body.Bytes(), &received)

		fmt.Printf("%v", received)

		assert.Equal(t, uuidToDelete, received.Uuid)
		assert.Equal(t, todoToCreate.Title, received.Title)
	})

	t.Run("Update todo", func(t *testing.T) {
		w := httptest.NewRecorder()

		// By reference, so yeah, we are modifying it directly
		todoToUpdate := todoToCreate
		todoToUpdate.Title = updatedTitle

		json_data, _ := json.Marshal(todoToUpdate)
		req, _ := http.NewRequest("PUT", "/todos/"+uuidToDelete, bytes.NewBuffer(json_data))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, 204, w.Code)
	})

	t.Run("Get todo", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todos/"+uuidToDelete, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var received models.Todo
		json.Unmarshal(w.Body.Bytes(), &received)

		fmt.Printf("%v", received)

		assert.Equal(t, uuidToDelete, received.Uuid)
		assert.Equal(t, updatedTitle, received.Title)
	})

	t.Run("Delete todo", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/todos/"+uuidToDelete, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 204, w.Code)
	})

	t.Run("List todos - should have zero now (deleted)", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todos", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var received []models.Todo
		json.Unmarshal(w.Body.Bytes(), &received)

		var expected []models.Todo
		assert.ElementsMatch(t, expected, received)
	})
}
