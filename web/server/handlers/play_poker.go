package handlers

import (
	"errors"
	"go-scrum-poker-bot/ui"
	"go-scrum-poker-bot/web/clients"
	"go-scrum-poker-bot/web/server/models"
	"net/http"

	"github.com/google/uuid"
)

func PlayPokerCommand(webClient clients.Client, uiBuilder *ui.Builder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.PostFormValue("channel_id") == "" || r.PostFormValue("text") == "" {
			w.Write(models.ResponseError(errors.New("Please write correct subject")))
			return
		}

		resp := webClient.Make(&clients.Request{
			URL:    "https://slack.com/api/chat.postMessage",
			Method: "POST",
			Json: uiBuilder.Build(
				r.PostFormValue("channel_id"),
				uuid.New().String(),
				r.PostFormValue("text"),
				nil,
				false,
			),
		})
		if resp.Error != nil {
			w.Write(models.ResponseError(resp.Error))
			return
		}
	})
}
