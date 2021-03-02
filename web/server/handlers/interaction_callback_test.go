package handlers_test

import (
	"encoding/json"
	"go-scrum-poker-bot/config"
	"go-scrum-poker-bot/ui"
	"go-scrum-poker-bot/web/server/handlers"
	"go-scrum-poker-bot/web/server/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInteractionCallbackHandlerActions(t *testing.T) {
	config := config.NewConfig()
	mockClient := &MockClient{}
	mockUserStorage := &MockUserStorage{}
	mockSessionStorage := &MockSessionStorage{}
	uiBuilder := ui.NewBuilder(config)

	router := http.NewServeMux()
	router.Handle(
		"/interactivity",
		handlers.InteractionCallback(mockUserStorage, mockSessionStorage, uiBuilder, mockClient),
	)

	actions := []*models.Action{
		{
			BlockID:        "test",
			ActionID:       ui.RESULTS_VISIBILITY_ACTION_ID,
			Value:          "test",
			SelectedOption: nil,
		},
		{
			BlockID:        "test",
			ActionID:       ui.VOTE_ACTION_ID,
			Value:          "test",
			SelectedOption: &models.SelectedOption{Value: "1"},
		},
	}

	for _, action := range actions {
		responseRec := httptest.NewRecorder()

		data, _ := json.Marshal(models.Callback{
			ResponseURL: "test",
			User:        &models.User{Username: "test"},
			Actions:     []*models.Action{action},
			Message: &models.Message{
				Blocks: []*models.Block{
					{
						Type:    "test",
						BlockID: ui.SUBJECT_BLOCK_ID,
						Text:    &models.Text{Type: "test", Text: "test"},
					},
				},
			},
		})
		payload := url.Values{"payload": {string(data)}}.Encode()
		request, err := http.NewRequest("POST", "/interactivity", strings.NewReader(payload))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		router.ServeHTTP(responseRec, request)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, responseRec.Code)
		assert.Empty(t, responseRec.Body.String())
		assert.Equal(t, true, mockClient.Called)
	}
}
