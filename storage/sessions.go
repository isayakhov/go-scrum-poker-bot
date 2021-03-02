package storage

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const SESSION_VOTES_HIDDEN_TPL = "SESSION:%s:VOTES_HIDDEN"

type SessionRedisStorage struct {
	redis   *redis.Client
	context context.Context
}

func NewSessionRedisStorage(redisClient *redis.Client) *SessionRedisStorage {
	return &SessionRedisStorage{
		redis:   redisClient,
		context: context.Background(),
	}
}

func (s *SessionRedisStorage) GetVisibility(sessionID string) bool {
	value, _ := strconv.ParseBool(
		s.redis.Get(s.context, fmt.Sprintf(SESSION_VOTES_HIDDEN_TPL, sessionID)).Val(),
	)

	return value
}

func (s *SessionRedisStorage) SetVisibility(sessionID string, state bool) error {
	return s.redis.Set(
		s.context,
		fmt.Sprintf(SESSION_VOTES_HIDDEN_TPL, sessionID),
		strconv.FormatBool(state),
		-1,
	).Err()
}
