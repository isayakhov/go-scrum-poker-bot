package blocks

type Action struct {
	Type     BlockType `json:"type"`
	BlockID  string    `json:"block_id"`
	Elements []Block   `json:"elements"`
}

func (a *Action) BlockType() BlockType {
	return ActionBlockType
}
