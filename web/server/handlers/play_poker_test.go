package handlers_test

import (
	"errors"
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

func TestPlayPokerHandler(t *testing.T) {
	config := config.NewConfig()
	mockClient := &MockClient{}
	uiBuilder := ui.NewBuilder(config)

	responseRec := httptest.NewRecorder()

	router := http.NewServeMux()
	router.Handle("/play-poker", handlers.PlayPokerCommand(mockClient, uiBuilder))

	payload := url.Values{"channel_id": {"test"}, "text": {"test"}}.Encode()
	request, err := http.NewRequest("POST", "/play-poker", strings.NewReader(payload))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(responseRec, request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, responseRec.Code)
	assert.Empty(t, responseRec.Body.String())
	assert.Equal(t, true, mockClient.Called)
}

func TestPlayPokerHandlerEmptyBodyError(t *testing.T) {
	config := config.NewConfig()
	mockClient := &MockClient{}
	uiBuilder := ui.NewBuilder(config)

	responseRec := httptest.NewRecorder()

	router := http.NewServeMux()
	router.Handle("/play-poker", handlers.PlayPokerCommand(mockClient, uiBuilder))

	payload := url.Values{}.Encode()
	request, _ := http.NewRequest("POST", "/play-poker", strings.NewReader(payload))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(responseRec, request)

	expected := string(models.ResponseError(errors.New("Please write correct subject")))

	assert.Equal(t, http.StatusOK, responseRec.Code)
	assert.Equal(t, expected, responseRec.Body.String())
	assert.Equal(t, false, mockClient.Called)
}

func TestPlayPokerHandlerRequestError(t *testing.T) {
	errMsg := "Error msg"
	config := config.NewConfig()
	mockClient := &MockClient{Error: errMsg}
	uiBuilder := ui.NewBuilder(config)

	responseRec := httptest.NewRecorder()

	router := http.NewServeMux()
	router.Handle("/play-poker", handlers.PlayPokerCommand(mockClient, uiBuilder))

	payload := url.Values{"channel_id": {"test"}, "text": {"test"}}.Encode()
	request, _ := http.NewRequest("POST", "/play-poker", strings.NewReader(payload))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(responseRec, request)

	expected := string(models.ResponseError(errors.New(errMsg)))

	assert.Equal(t, http.StatusOK, responseRec.Code)
	assert.Equal(t, expected, responseRec.Body.String())
	assert.Equal(t, true, mockClient.Called)
}
