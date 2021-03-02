package blocks

type Select struct {
	Type        BlockType `json:"type"`
	Placeholder *Text     `json:"placeholder"`
	Options     []*Option `json:"options"`
	ActionID    string    `json:"action_id"`
}

func (b *Select) BlockType() BlockType {
	return StaticSelectBlockType
}

type Option struct {
	Text  *Text  `json:"text"`
	Value string `json:"value"`
}
