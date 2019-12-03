package smyth

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type Attribute struct {
	Name         string   `json:"name"`
	Common       TypeList `json:"common"`
	Uncommon     TypeList `json:"uncommon"`
	Rare         TypeList `json:"rare"`
	WeightFactor float64  `json:"weight_factor"`
	PrefixNames  []string `json:"prefix_names"`
	SuffixNames  []string `json:"suffix_names"`
}

// readAttributes reads the JSON-encoded Attributes from the provided Reader.
func readAttributes(r io.Reader) ([]Attribute, error) {
	var attr []Attribute

	if err := json.NewDecoder(r).Decode(&attr); err != nil {
		return nil, errors.Wrap(err, "cannot decode Attribute from io.Reader")
	}

	for _, a := range attr {
		if a.WeightFactor <= 0 {
			a.WeightFactor = 1
		}

		if a.Common.ValueFactor <= 0 {
			a.Common.ValueFactor = 1
		}

		if a.Uncommon.ValueFactor <= 0 {
			a.Uncommon.ValueFactor = 1
		}

		if a.Rare.ValueFactor <= 0 {
			a.Rare.ValueFactor = 1
		}
	}

	return attr, nil
}
