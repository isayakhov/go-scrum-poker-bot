package clients

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler func(request *Request) *Response
type Middleware func(handler Handler, request *Request) Handler

type Client interface {
	Make(request *Request) *Response
}

type BasicClient struct {
	client     *http.Client
	middleware []Middleware
}

func NewBasicClient(client *http.Client, middleware []Middleware) Client {
	return &BasicClient{client: client, middleware: middleware}
}

func (c *BasicClient) makeRequest(request *Request) *Response {
	payload, err := request.ToBytes()
	if err != nil {
		return &Response{Error: err}
	}

	req, err := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(payload))
	if err != nil {
		return &Response{Error: err}
	}

	for name, value := range request.Headers {
		req.Header.Add(name, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return &Response{Error: err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{Error: err}
	}

	err = nil
	if resp.StatusCode > http.StatusIMUsed || resp.StatusCode < http.StatusOK {
		err = fmt.Errorf("Bad response. Status: %d, Body: %s", resp.StatusCode, string(body))
	}

	return &Response{
		Status:  resp.StatusCode,
		Body:    body,
		Headers: resp.Header,
		Error:   err,
	}
}

func (c *BasicClient) Make(request *Request) *Response {
	if request.Headers == nil {
		request.Headers = make(map[string]string)
	}
	handler := c.makeRequest
	for _, middleware := range c.middleware {
		handler = middleware(handler, request)
	}

	return handler(request)
}
