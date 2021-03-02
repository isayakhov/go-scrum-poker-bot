package blocks

type Button struct {
	Type     BlockType `json:"type"`
	Text     *Text     `json:"text"`
	Style    string    `json:"style,omitempty"`
	Value    string    `json:"value"`
	ActionID string    `json:"action_id"`
}

func (b *Button) BlockType() BlockType {
	return ButtonBlockType
}
