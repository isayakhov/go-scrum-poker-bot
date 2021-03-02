package handlers

import (
	"encoding/json"
	"errors"
	"go-scrum-poker-bot/web/server/models"
	"net/http"
)

func Healthcheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(models.Healthcheck{Status: "OK"})
		if err != nil {
			w.Write(models.ResponseError(errors.New("Something worng")))
			return
		}
		w.Write(data)
	})
}
