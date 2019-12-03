package smyth

type Resources struct {
	Recipes    map[string]Recipe
	Attributes map[string]Attribute
	AttrGroups map[string]AttrGroup
}
