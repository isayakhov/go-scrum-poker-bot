package middleware_test

import (
	"bytes"
	"errors"
	"go-scrum-poker-bot/web/server/middleware"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type panicHandler struct {
	err interface{}
}

func (h *panicHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic(h.err)
}

func TestRecoverMiddleware(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	logger.SetOutput(&buf)

	for _, err := range []interface{}{errors.New("String"), "string", 1} {
		handler := &panicHandler{err: err}
		responseRec := httptest.NewRecorder()

		router := http.NewServeMux()
		router.Handle("/test", middleware.Recover(logger)(handler))

		request, err := http.NewRequest("GET", "/test", strings.NewReader(""))

		router.ServeHTTP(responseRec, request)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, responseRec.Code)
		assert.NotEmpty(t, buf.String())
	}
}
