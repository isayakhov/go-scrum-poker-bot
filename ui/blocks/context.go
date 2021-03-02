package blocks

type Context struct {
	BlockID  string    `json:"block_id,omitempty"`
	Type     BlockType `json:"type"`
	Elements []*Text   `json:"elements"`
}

func (c *Context) BlockType() BlockType {
	return ContextBlockType
}
