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

func TestTodosErrors(t *testing.T) {
	// Setup
	router := api.SetupApi()

	t.Run("Create a new todo - should fail, missing title", func(t *testing.T) {
		todoToCreate := models.Todo{Uuid: uuid.New().String(), OwnerUuid: uuid.New().String(), State: "ACTIVE"}

		w := httptest.NewRecorder()
		json_data, _ := json.Marshal(todoToCreate)
		req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(json_data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("Create a new todo - should fail, empty title", func(t *testing.T) {
		todoToCreate := models.Todo{Uuid: uuid.New().String(), Title: "", OwnerUuid: uuid.New().String(), State: "ACTIVE"}

		w := httptest.NewRecorder()
		json_data, _ := json.Marshal(todoToCreate)
		req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(json_data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("Create a new todo - should fail, wrong state", func(t *testing.T) {
		todoToCreate := models.Todo{Uuid: uuid.New().String(), Title: "", OwnerUuid: uuid.New().String(), State: "RANDOM_WORD"}

		w := httptest.NewRecorder()
		json_data, _ := json.Marshal(todoToCreate)
		req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(json_data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("Updating a todo - should fail, todo doesn't exist", func(t *testing.T) {
		var uuidToUpdate = uuid.New().String()
		todoToUpdate := models.Todo{Uuid: uuidToUpdate, Title: "Do the groceries", OwnerUuid: uuid.New().String(), State: "ACTIVE"}

		w := httptest.NewRecorder()
		json_data, _ := json.Marshal(todoToUpdate)
		req, _ := http.NewRequest("PUT", "/todos/"+uuidToUpdate, bytes.NewBuffer(json_data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})
}
