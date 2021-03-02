package handlers

import (
	"go-scrum-poker-bot/storage"
	"go-scrum-poker-bot/ui"
	"go-scrum-poker-bot/ui/blocks"
	"go-scrum-poker-bot/web/clients"
	"go-scrum-poker-bot/web/server/models"
	"net/http"
)

func InteractionCallback(
	userStorage storage.UserStorage,
	sessionStorage storage.SessionStorage,
	uiBuilder *ui.Builder,
	webClient clients.Client,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var callback models.Callback
		data, err := callback.SerializedData([]byte(r.PostFormValue("payload")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users := userStorage.All(data.SessionID)
		visible := sessionStorage.GetVisibility(data.SessionID)

		err = nil
		switch data.Action.ActionID {
		case ui.VOTE_ACTION_ID:
			users[callback.User.Username] = data.Action.SelectedOption.Value
			err = userStorage.Save(data.SessionID, callback.User.Username, data.Action.SelectedOption.Value)
		case ui.RESULTS_VISIBILITY_ACTION_ID:
			visible = !visible
			err = sessionStorage.SetVisibility(data.SessionID, visible)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := webClient.Make(&clients.Request{
			URL:    callback.ResponseURL,
			Method: "POST",
			Json: &blocks.Interactive{
				ReplaceOriginal: true,
				Blocks:          uiBuilder.BuildBlocks(data.Subject, users, data.SessionID, visible),
				LinkNames:       true,
			},
		})
		if resp.Error != nil {
			http.Error(w, resp.Error.Error(), http.StatusInternalServerError)
			return
		}
	})
}
