package clients_test

import (
	"bytes"
	"go-scrum-poker-bot/web/clients"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RoundTripFunc func(request *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestMakeRequest(t *testing.T) {
	url := "https://example.com/ok"

	httpClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), url)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("OK")),
			Header:     make(http.Header),
		}
	})

	webClient := clients.NewBasicClient(httpClient, nil)
	response := webClient.Make(&clients.Request{
		URL:     url,
		Method:  "GET",
		Headers: map[string]string{"Content-Type": "application/json"},
		Json:    nil,
	})

	assert.Equal(t, http.StatusOK, response.Status)
}

func TestMakeRequestError(t *testing.T) {
	url := "https://example.com/error"

	httpClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), url)

		return &http.Response{
			StatusCode: http.StatusBadGateway,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Bad gateway")),
			Header:     make(http.Header),
		}
	})

	webClient := clients.NewBasicClient(httpClient, nil)
	response := webClient.Make(&clients.Request{
		URL:     url,
		Method:  "GET",
		Headers: map[string]string{"Content-Type": "application/json"},
		Json:    nil,
	})

	assert.Equal(t, http.StatusBadGateway, response.Status)
}
