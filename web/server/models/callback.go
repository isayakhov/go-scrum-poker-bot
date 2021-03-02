package models

import (
	"encoding/json"
	"errors"
	"go-scrum-poker-bot/ui"
)

type User struct {
	Username string `json:"username"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Block struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id"`
	Text    *Text  `json:"text,omitempty"`
}

type Message struct {
	Blocks []*Block `json:"blocks,omitempty"`
}

type SelectedOption struct {
	Value string `json:"value"`
}

type Action struct {
	BlockID        string          `json:"block_id"`
	ActionID       string          `json:"action_id"`
	Value          string          `json:"value,omitempty"`
	SelectedOption *SelectedOption `json:"selected_option,omitempty"`
}

type SerializedData struct {
	SessionID string
	Subject   string
	Action    *Action
}

type Callback struct {
	ResponseURL string    `json:"response_url"`
	User        *User     `json:"user"`
	Actions     []*Action `json:"actions"`
	Message     *Message  `json:"message,omitempty"`
}

func (c *Callback) getSessionID() (string, error) {
	for _, action := range c.Actions {
		if action.BlockID != "" {
			return action.BlockID, nil
		}
	}

	return "", errors.New("Invalid session ID")
}

func (c *Callback) getSubject() (string, error) {
	for _, block := range c.Message.Blocks {
		if block.BlockID == ui.SUBJECT_BLOCK_ID && block.Text != nil {
			return block.Text.Text, nil
		}
	}

	return "", errors.New("Invalid subject")
}

func (c *Callback) getAction() (*Action, error) {
	for _, action := range c.Actions {
		if action.ActionID == ui.VOTE_ACTION_ID || action.ActionID == ui.RESULTS_VISIBILITY_ACTION_ID {
			return action, nil
		}
	}

	return nil, errors.New("Invalid action")
}

func (c *Callback) SerializedData(data []byte) (*SerializedData, error) {
	err := json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	sessionID, err := c.getSessionID()
	if err != nil {
		return nil, err
	}

	subject, err := c.getSubject()
	if err != nil {
		return nil, err
	}

	action, err := c.getAction()
	if err != nil {
		return nil, err
	}

	return &SerializedData{
		SessionID: sessionID,
		Subject:   subject,
		Action:    action,
	}, nil
}
