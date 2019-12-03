package smyth

type Property struct {
	Name           string   `json:"name"`
	Required       bool     `json:"required"`
	AttributeNames []string `json:"attribute_names"`
	AttrGroupNames []string `json:"attribute_group_names"`
}
