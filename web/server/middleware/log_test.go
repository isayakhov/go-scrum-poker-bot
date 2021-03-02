package middleware_test

import (
	"bytes"
	"go-scrum-poker-bot/web/server/middleware"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type logHandler struct{}

func (h *logHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func TestLogMiddleware(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	logger.SetOutput(&buf)

	handler := &logHandler{}
	responseRec := httptest.NewRecorder()

	router := http.NewServeMux()
	router.Handle("/test", middleware.Log(logger)(handler))

	request, err := http.NewRequest("GET", "/test", strings.NewReader(""))

	router.ServeHTTP(responseRec, request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, responseRec.Code)
	assert.NotEmpty(t, buf.String())
}
