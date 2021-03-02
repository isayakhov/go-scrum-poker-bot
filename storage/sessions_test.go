package storage_test

import (
	"errors"
	"fmt"
	"go-scrum-poker-bot/storage"
	"strconv"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestGetVisibility(t *testing.T) {
	sessionID, state := "test", true

	redisClient, mock := redismock.NewClientMock()

	mock.ExpectGet(
		fmt.Sprintf(storage.SESSION_VOTES_HIDDEN_TPL, sessionID),
	).SetVal(strconv.FormatBool(state))

	sessionStorage := storage.NewSessionRedisStorage(redisClient)

	assert.Equal(t, state, sessionStorage.GetVisibility(sessionID))
}

func TestSetVisibility(t *testing.T) {
	sessionID, state := "test", true

	redisClient, mock := redismock.NewClientMock()

	mock.ExpectSet(
		fmt.Sprintf(storage.SESSION_VOTES_HIDDEN_TPL, sessionID),
		strconv.FormatBool(state),
		-1,
	).SetVal("1")

	sessionStorage := storage.NewSessionRedisStorage(redisClient)

	assert.Equal(t, nil, sessionStorage.SetVisibility(sessionID, state))
}

func TestSetVisibilityErr(t *testing.T) {
	sessionID, state, err := "test", true, errors.New("ERROR")

	redisClient, mock := redismock.NewClientMock()

	mock.ExpectSet(
		fmt.Sprintf(storage.SESSION_VOTES_HIDDEN_TPL, sessionID),
		strconv.FormatBool(state),
		-1,
	).SetErr(err)

	sessionStorage := storage.NewSessionRedisStorage(redisClient)

	assert.Equal(t, err, sessionStorage.SetVisibility(sessionID, state))
}
