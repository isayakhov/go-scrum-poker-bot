package clients_test

import (
	"errors"
	"go-scrum-poker-bot/web/clients"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseJson(t *testing.T) {
	to := struct {
		TestKey string `json:"test_key"`
	}{}
	response := clients.Response{
		Status:  200,
		Headers: nil,
		Body:    []byte(`{"test_key": "test_value"}`),
		Error:   nil,
	}

	err := response.Json(&to)

	assert.Equal(t, nil, err)
	assert.Equal(t, "test_value", to.TestKey)
}

func TestResponseJsonError(t *testing.T) {
	expectedErr := errors.New("Error!")
	response := clients.Response{
		Status:  200,
		Headers: nil,
		Body:    nil,
		Error:   expectedErr,
	}

	err := response.Json(map[string]string{})

	assert.Equal(t, expectedErr, err)
}
