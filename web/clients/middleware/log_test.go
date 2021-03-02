package middleware_test

import (
	"bytes"
	"go-scrum-poker-bot/web/clients"
	"go-scrum-poker-bot/web/clients/middleware"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogMiddleware(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	logger.SetOutput(&buf)

	request := &clients.Request{
		Headers: map[string]string{},
	}
	handler := middleware.Log(logger)(
		func(request *clients.Request) *clients.Response {
			return &clients.Response{}
		},
		request,
	)
	handler(request)

	assert.NotEmpty(t, buf.String())
}
