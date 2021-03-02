package blocks

type Interactive struct {
	ReplaceOriginal bool    `json:"replace_original"`
	Blocks          []Block `json:"blocks"`
	LinkNames       bool    `json:"link_names"`
}
