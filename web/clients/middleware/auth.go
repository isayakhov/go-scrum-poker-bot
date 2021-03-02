package middleware

import (
	"fmt"
	"go-scrum-poker-bot/web/clients"
)

func Auth(token string) clients.Middleware {
	return func(handler clients.Handler, request *clients.Request) clients.Handler {
		return func(request *clients.Request) *clients.Response {
			request.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
			return handler(request)
		}
	}
}
