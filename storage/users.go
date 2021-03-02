package storage

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const SESSION_USERS_TPL = "SESSION:%s:USERS"
const USER_VOTE_TPL = "SESSION:%s:USERNAME:%s:VOTE"

type UserRedisStorage struct {
	redis   *redis.Client
	context context.Context
}

func NewUserRedisStorage(redisClient *redis.Client) *UserRedisStorage {
	return &UserRedisStorage{
		redis:   redisClient,
		context: context.Background(),
	}
}

func (s *UserRedisStorage) All(sessionID string) map[string]string {
	users := make(map[string]string)

	for _, username := range s.redis.SMembers(s.context, fmt.Sprintf(SESSION_USERS_TPL, sessionID)).Val() {
		users[username] = s.redis.Get(s.context, fmt.Sprintf(USER_VOTE_TPL, sessionID, username)).Val()
	}
	return users
}

func (s *UserRedisStorage) Save(sessionID string, username string, value string) error {
	err := s.redis.SAdd(s.context, fmt.Sprintf(SESSION_USERS_TPL, sessionID), username).Err()
	if err != nil {
		return err
	}

	err = s.redis.Set(s.context, fmt.Sprintf(USER_VOTE_TPL, sessionID, username), value, -1).Err()
	if err != nil {
		return err
	}

	return nil
}
