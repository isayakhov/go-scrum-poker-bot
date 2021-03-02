package handlers_test

import (
	"errors"
	"go-scrum-poker-bot/web/clients"
)

type MockClient struct {
	Called bool
	Error  string
}

func (c *MockClient) Make(request *clients.Request) *clients.Response {
	c.Called = true

	var err error = nil
	if c.Error != "" {
		err = errors.New(c.Error)
	}
	return &clients.Response{Error: err}
}

type MockUserStorage struct{}

func (s *MockUserStorage) All(sessionID string) map[string]string {
	return map[string]string{"user": "1"}
}

func (s *MockUserStorage) Save(sessionID string, username string, value string) error {
	return nil
}

type MockSessionStorage struct{}

func (s *MockSessionStorage) GetVisibility(sessionID string) bool {
	return true
}

func (s *MockSessionStorage) SetVisibility(sessionID string, state bool) error {
	return nil
}
