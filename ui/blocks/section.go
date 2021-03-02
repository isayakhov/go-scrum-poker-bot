package blocks

type Section struct {
	BlockID string    `json:"block_id,omitempty"`
	Type    BlockType `json:"type"`
	Text    *Text     `json:"text"`
}

func (s *Section) BlockType() BlockType {
	return SectionBlockType
}
