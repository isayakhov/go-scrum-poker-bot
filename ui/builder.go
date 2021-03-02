package ui

import (
	"fmt"
	"go-scrum-poker-bot/config"
	"go-scrum-poker-bot/ui/blocks"
	"strings"
)

type Builder struct {
	config *config.Config
}

func NewBuilder(config *config.Config) *Builder {
	return &Builder{config: config}
}

func (b *Builder) getGetResultsText(visible bool) string {
	if visible {
		return VOTES_VISIBLE_STATE
	} else {
		return VOTES_INVISIBLE_STATE
	}
}

func (b *Builder) getResults(users map[string]string, visible bool) string {
	if len(users) == 0 {
		return NOBODY_VOTED
	}

	var votes []string
	for user, vote := range users {
		if !visible {
			vote = VOTE_HIDDEN_VALUE
		}
		votes = append(votes, fmt.Sprintf(USER_VOTE, user, vote))
	}
	return strings.Join(votes, ", ")
}

func (b *Builder) getOptions() []*blocks.Option {
	var options []*blocks.Option
	for _, value := range b.config.App.PokerRanks {
		options = append(options, &blocks.Option{Text: &blocks.Text{Type: TYPE_PLAIN_TEXT, Text: value}, Value: value})
	}

	return options
}

func (b *Builder) BuildBlocks(
	subject string,
	users map[string]string,
	sessionID string,
	visible bool,
) []blocks.Block {
	return []blocks.Block{
		&blocks.Section{
			BlockID: SUBJECT_BLOCK_ID,
			Type:    blocks.SectionBlockType,
			Text: &blocks.Text{
				Type: TYPE_MARKDOWN,
				Text: subject,
			},
		},
		&blocks.Context{
			BlockID: USERS_BLOCK_ID,
			Type:    blocks.ContextBlockType,
			Elements: []*blocks.Text{
				{Type: TYPE_MARKDOWN, Text: b.getResults(users, visible)},
			},
		},
		&blocks.Action{
			Type:    TYPE_ACTIONS,
			BlockID: sessionID,
			Elements: []blocks.Block{
				&blocks.Select{
					ActionID: VOTE_ACTION_ID,
					Type:     blocks.StaticSelectBlockType,
					Placeholder: &blocks.Text{
						Type: TYPE_PLAIN_TEXT,
						Text: CHOOSE_YOUR_VOTE_TEXT,
					},
					Options: b.getOptions(),
				},
				&blocks.Button{
					ActionID: RESULTS_VISIBILITY_ACTION_ID,
					Type:     blocks.ButtonBlockType,
					Text: &blocks.Text{
						Type: TYPE_PLAIN_TEXT,
						Text: b.getGetResultsText(visible),
					},
					Value: RESULTS_VISIBILITY_VALUE,
					Style: STYLE_PRIMARY,
				},
			},
		},
	}
}

func (b *Builder) Build(
	channel string,
	sessionID string,
	subject string,
	users map[string]string,
	visible bool,
) *blocks.Blocks {
	return &blocks.Blocks{
		Channel: channel,
		Blocks:  b.BuildBlocks(subject, users, sessionID, visible),
	}
}
