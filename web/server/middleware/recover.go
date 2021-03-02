package middleware

import (
	"errors"
	"log"
	"net/http"
)

func Recover(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				var err error
				rec := recover()
				if rec != nil {
					switch t := rec.(type) {
					case string:
						err = errors.New(t)
					case error:
						err = t
					default:
						err = errors.New("Unknown error")
					}
					logger.Printf(
						"PANIC: [%s]: %s - %s - %s. Error: %s. Trying to recover state...",
						r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), err.Error(),
					)
					http.Error(w, "Something went wrong", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
