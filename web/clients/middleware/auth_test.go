package middleware_test

import (
	"fmt"
	"go-scrum-poker-bot/web/clients"
	"go-scrum-poker-bot/web/clients/middleware"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	token := "test"
	request := &clients.Request{
		Headers: map[string]string{},
	}
	handler := middleware.Auth(token)(
		func(request *clients.Request) *clients.Response {
			return &clients.Response{}
		},
		request,
	)
	handler(request)

	assert.Equal(t, map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)}, request.Headers)
}
