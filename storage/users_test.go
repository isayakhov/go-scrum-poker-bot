package storage_test

import (
	"errors"
	"fmt"
	"go-scrum-poker-bot/storage"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	sessionID, username, value := "test", "user", "1"

	redisClient, mock := redismock.NewClientMock()
	usersStorage := storage.NewUserRedisStorage(redisClient)

	mock.ExpectSMembers(
		fmt.Sprintf(storage.SESSION_USERS_TPL, sessionID),
	).SetVal([]string{username})
	mock.ExpectGet(
		fmt.Sprintf(storage.USER_VOTE_TPL, sessionID, username),
	).SetVal(value)

	assert.Equal(t, map[string]string{username: value}, usersStorage.All(sessionID))
}

func TestSave(t *testing.T) {
	sessionID, username, value := "test", "user", "1"

	redisClient, mock := redismock.NewClientMock()
	usersStorage := storage.NewUserRedisStorage(redisClient)

	mock.ExpectSAdd(
		fmt.Sprintf(storage.SESSION_USERS_TPL, sessionID),
		username,
	).SetVal(1)
	mock.ExpectSet(
		fmt.Sprintf(storage.USER_VOTE_TPL, sessionID, username),
		value,
		-1,
	).SetVal(value)

	assert.Equal(t, nil, usersStorage.Save(sessionID, username, value))
}

func TestSaveSAddErr(t *testing.T) {
	sessionID, username, value, err := "test", "user", "1", errors.New("ERROR")

	redisClient, mock := redismock.NewClientMock()
	usersStorage := storage.NewUserRedisStorage(redisClient)

	mock.ExpectSAdd(
		fmt.Sprintf(storage.SESSION_USERS_TPL, sessionID),
		username,
	).SetErr(err)

	assert.Equal(t, err, usersStorage.Save(sessionID, username, value))
}

func TestSaveSetErr(t *testing.T) {
	sessionID, username, value, err := "test", "user", "1", errors.New("ERROR")

	redisClient, mock := redismock.NewClientMock()
	usersStorage := storage.NewUserRedisStorage(redisClient)

	mock.ExpectSAdd(
		fmt.Sprintf(storage.SESSION_USERS_TPL, sessionID),
		username,
	).SetVal(1)
	mock.ExpectSet(
		fmt.Sprintf(storage.USER_VOTE_TPL, sessionID, username),
		value,
		-1,
	).SetErr(err)

	assert.Equal(t, err, usersStorage.Save(sessionID, username, value))
}
