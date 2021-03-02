package middleware

import "go-scrum-poker-bot/web/clients"

func JsonContentType(handler clients.Handler, request *clients.Request) clients.Handler {
	return func(request *clients.Request) *clients.Response {
		request.Headers["Content-Type"] = "application/json"
		return handler(request)
	}
}
