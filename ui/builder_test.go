package ui_test

import (
	"fmt"
	"go-scrum-poker-bot/config"
	"go-scrum-poker-bot/ui"
	"go-scrum-poker-bot/ui/blocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilderBuild(t *testing.T) {
	config := config.NewConfig()
	builder := ui.NewBuilder(config)

	type data struct {
		Channel        string
		SessionID      string
		Subject        string
		Users          map[string]string
		UsersVotesText string
		Visible        bool
		VisibleText    string
	}

	testData := []data{
		{
			Channel:        "test_channel1",
			SessionID:      "test_session1",
			Subject:        "test_subject1",
			Users:          nil,
			UsersVotesText: ui.NOBODY_VOTED,
			Visible:        true,
			VisibleText:    ui.VOTES_VISIBLE_STATE,
		},
		{
			Channel:        "test_channel2",
			SessionID:      "test_session2",
			Subject:        "test_subject2",
			Users:          map[string]string{"test": "1"},
			UsersVotesText: fmt.Sprintf("@test: *%s*", ui.VOTE_HIDDEN_VALUE),
			Visible:        false,
			VisibleText:    ui.VOTES_INVISIBLE_STATE,
		},
	}

	for _, expected := range testData {
		builtUI := builder.Build(
			expected.Channel,
			expected.SessionID,
			expected.Subject,
			expected.Users,
			expected.Visible,
		)

		actualChannel := builtUI.Channel
		actualSubject := builtUI.Blocks[0].(*blocks.Section).Text.Text
		actualVotesText := builtUI.Blocks[1].(*blocks.Context).Elements[0].Text
		actualRankValue := builtUI.Blocks[2].(*blocks.Action).Elements[0].(*blocks.Select).Options[0].Text.Text
		actualVisibleText := builtUI.Blocks[2].(*blocks.Action).Elements[1].(*blocks.Button).Text.Text

		assert.Equal(t, expected.Channel, actualChannel)
		assert.Equal(t, expected.Subject, actualSubject)
		assert.Equal(t, expected.UsersVotesText, actualVotesText)
		assert.Equal(t, config.App.PokerRanks[0], actualRankValue)
		assert.Equal(t, expected.VisibleText, actualVisibleText)
	}
}
