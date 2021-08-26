package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GuillaumeDeconinck/todos-go/internal/api"
	"github.com/stretchr/testify/assert"
)

type PingMessage struct {
	Message string `json:"message"`
}

func TestHandlePing(t *testing.T) {
	router := api.SetupApi()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var p PingMessage
	json.Unmarshal(w.Body.Bytes(), &p)
	assert.Equal(t, PingMessage{Message: "pong"}, p)
}
