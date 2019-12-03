package smyth

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type Recipe struct {
	Name       string      `json:"name"`
	BaseValue  float64     `json:"base_value"`
	BaseWeight float64     `json:"base_weight"`
	Comps      []Component `json:"components"`
}

// ReadRecipe reads the JSON-encoded Recipes from the provided Reader.
func readRecipes(r io.Reader) ([]Recipe, error) {
	var rec []Recipe

	if err := json.NewDecoder(r).Decode(&rec); err != nil {
		return nil, errors.Wrap(err, "cannot decode Recipe from io.Reader")
	}

	return rec, nil
}
