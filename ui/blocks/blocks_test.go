package blocks_test

import (
	"go-scrum-poker-bot/ui/blocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilderBlockType(t *testing.T) {
	blocks := map[blocks.BlockType]blocks.Block{
		blocks.ActionBlockType:       &blocks.Action{},
		blocks.ButtonBlockType:       &blocks.Button{},
		blocks.ContextBlockType:      &blocks.Context{},
		blocks.StaticSelectBlockType: &blocks.Select{},
		blocks.SectionBlockType:      &blocks.Section{},
	}
	for blockType, block := range blocks {
		assert.Equal(t, blockType, block.BlockType())
	}
}
