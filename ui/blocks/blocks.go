package blocks

type BlockType string

const (
	ContextBlockType      BlockType = "context"
	SectionBlockType      BlockType = "section"
	StaticSelectBlockType BlockType = "static_select"
	ActionBlockType       BlockType = "action"
	ButtonBlockType       BlockType = "button"
)

type Block interface {
	BlockType() BlockType
}

type Blocks struct {
	Channel string  `json:"channel"`
	Blocks  []Block `json:"blocks"`
}
