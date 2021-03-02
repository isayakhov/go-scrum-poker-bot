package middleware

import (
	"log"
	"net/http"
)

func Log(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				logger.Printf(
					"Handle request: [%s]: %s - %s - %s",
					r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(),
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
