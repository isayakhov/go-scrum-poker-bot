package models

import (
	"encoding/json"
	"fmt"
)

type SlackError struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func ResponseError(err error) []byte {
	resp, err := json.Marshal(
		SlackError{
			ResponseType: "ephemeral",
			Text:         fmt.Sprintf("Sorry, there is some error happened. Error: %s", err.Error()),
		},
	)
	if err != nil {
		return []byte("Sorry. Some error happened")
	}
	return resp
}
