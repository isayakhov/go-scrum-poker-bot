package handlers_test

import (
	"go-scrum-poker-bot/web/server/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheckHandler(t *testing.T) {
	responseRec := httptest.NewRecorder()

	router := http.NewServeMux()
	router.Handle("/healthcheck", handlers.Healthcheck())

	request, err := http.NewRequest("POST", "/healthcheck", strings.NewReader(""))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(responseRec, request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, responseRec.Code)
	assert.Equal(t, `{"status":"OK"}`, responseRec.Body.String())
}
