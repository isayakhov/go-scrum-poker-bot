package clients_test

import (
	"encoding/json"
	"go-scrum-poker-bot/web/clients"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestToBytes(t *testing.T) {
	testCases := []struct {
		json interface{}
		data []byte
		err  error
	}{
		{map[string]string{"test_key": "test_value"}, []byte("{\"test_key\":\"test_value\"}"), nil},
		{nil, []byte{}, nil},
		{make(chan int), []byte{}, &json.UnsupportedTypeError{Type: reflect.TypeOf(make(chan int))}},
	}

	for _, testCase := range testCases {
		request := clients.Request{
			URL:     "https://example.com",
			Method:  "GET",
			Headers: nil,
			Json:    testCase.json,
		}

		actual, err := request.ToBytes()

		assert.Equal(t, testCase.err, err)
		assert.Equal(t, testCase.data, actual)
	}
}
