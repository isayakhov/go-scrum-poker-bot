package models_test

import (
	"encoding/json"
	"go-scrum-poker-bot/ui"
	"go-scrum-poker-bot/web/server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializedData(t *testing.T) {
	callback := models.Callback{
		ResponseURL: "test",
		User:        &models.User{Username: "test"},
		Actions: []*models.Action{
			{
				BlockID:        "test",
				ActionID:       ui.RESULTS_VISIBILITY_ACTION_ID,
				Value:          "test",
				SelectedOption: nil,
			},
		},
		Message: &models.Message{
			Blocks: []*models.Block{
				{
					Type:    "test",
					BlockID: ui.SUBJECT_BLOCK_ID,
					Text:    &models.Text{Type: "test", Text: "test"},
				},
			},
		},
	}
	data, _ := json.Marshal(callback)
	serializedData, err := callback.SerializedData(data)

	assert.Nil(t, err)
	assert.NotNil(t, serializedData)
}

func TestSerializedDataInvalidDataError(t *testing.T) {
	testCases := []struct {
		actionID  string
		sessionID string
		subjectID string
		errMsg    string
	}{
		{ui.RESULTS_VISIBILITY_ACTION_ID, "test", "wrong", "Invalid subject"},
		{ui.RESULTS_VISIBILITY_ACTION_ID, "", ui.SUBJECT_BLOCK_ID, "Invalid session ID"},
		{"wrong", "test", ui.SUBJECT_BLOCK_ID, "Invalid action"},
	}
	for _, testCase := range testCases {
		callback := models.Callback{
			ResponseURL: "test",
			User:        &models.User{Username: "test"},
			Actions: []*models.Action{
				{
					BlockID:        testCase.sessionID,
					ActionID:       testCase.actionID,
					Value:          "test",
					SelectedOption: nil,
				},
			},
			Message: &models.Message{
				Blocks: []*models.Block{
					{
						Type:    "test",
						BlockID: testCase.subjectID,
						Text:    &models.Text{Type: "test", Text: "test"},
					},
				},
			},
		}
		data, _ := json.Marshal(callback)
		serializedData, err := callback.SerializedData(data)

		assert.NotNil(t, err)
		assert.Nil(t, serializedData)
		assert.Equal(t, testCase.errMsg, err.Error())
	}
}
