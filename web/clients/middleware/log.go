package middleware

import (
	"go-scrum-poker-bot/web/clients"
	"log"
)

func Log(logger *log.Logger) clients.Middleware {
	return func(handler clients.Handler, request *clients.Request) clients.Handler {
		return func(request *clients.Request) *clients.Response {
			logger.Printf("Process request: [%s]: %s", request.Method, request.URL)

			response := handler(request)

			if response.Error != nil {
				logger.Printf("Error in request: [%s]: %s: %s", request.Method, request.URL, response.Error.Error())
			}

			logger.Printf("Finished request: [%s]: %s", request.Method, request.URL)

			return response
		}
	}
}
