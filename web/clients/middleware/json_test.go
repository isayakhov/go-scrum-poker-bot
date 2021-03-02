package middleware_test

import (
	"go-scrum-poker-bot/web/clients"
	"go-scrum-poker-bot/web/clients/middleware"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonMiddleware(t *testing.T) {
	request := &clients.Request{
		Headers: map[string]string{},
	}
	handler := middleware.JsonContentType(
		func(request *clients.Request) *clients.Response {
			return &clients.Response{}
		},
		request,
	)
	handler(request)

	assert.Equal(t, map[string]string{"Content-Type": "application/json"}, request.Headers)
}
