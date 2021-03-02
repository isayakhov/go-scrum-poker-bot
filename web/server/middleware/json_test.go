package middleware_test

import (
	"go-scrum-poker-bot/web/server/middleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type okHandler struct{}

func (h *okHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func TestJsonMiddleware(t *testing.T) {
	handler := &okHandler{}
	responseRec := httptest.NewRecorder()

	router := http.NewServeMux()
	router.Handle("/test", middleware.Json(handler))

	request, err := http.NewRequest("GET", "/test", strings.NewReader(""))

	router.ServeHTTP(responseRec, request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, responseRec.Code)
	assert.Equal(t, "application/json", responseRec.Header().Get("Content-Type"))
}
