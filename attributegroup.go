package smyth

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type AttrGroup struct {
	Name           string   `json:"name"`
	AttributeNames []string `json:"attribute_names"`
}

// readAttrGroups reads the JSON-encoded AttributeGroups from the provided
// Reader.
func readAttrGroups(r io.Reader) ([]AttrGroup, error) {
	var attr []AttrGroup

	if err := json.NewDecoder(r).Decode(&attr); err != nil {
		return nil, errors.Wrap(err, "cannot decode AttrGroup from io.Reader")
	}

	return attr, nil
}