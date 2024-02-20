package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Albums(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	body, _ := io.ReadAll(w.Body)
	var albums []Album
	json.Unmarshal(body, &albums)
	assert.Equal(t, 3, len(albums))
}
