package smyth

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/pkg/errors"
	"io"
)

var (
	ErrSetterEmpty = errors.New("setters require a slice with at least 1 element")
	ErrNameBlank   = errors.New("setters require struct with valid 'Name' field")
)

type Builder struct {
	src Resources
}

// NewBuilder returns an initialized Builder.
func NewBuilder() Builder {
	return Builder{
		src: Resources{
			Recipes:    make(map[string]Recipe),
			Attributes: make(map[string]Attribute),
			AttrGroups: make(map[string]AttrGroup),
		},
	}
}

// SetRecipes reads the JSON-encoded Recipes from the provided Reader and
// adds them to the receiver's Recipe map.
func (b *Builder) SetRecipes(r io.Reader) error {
	rec, err := readRecipes(r)
	if err != nil {
		return errors.Wrap(err, "cannot read Recipes")
	}

	if err := b.setRecipes(rec); err != nil {
		return errors.Wrap(err, "cannot set Recipes")
	}

	return nil
}

// setRecipes adds the provided Recipes to the receiver's Recipe map.
func (b *Builder) setRecipes(rec []Recipe) error {
	if len(rec) <= 0 {
		return ErrSetterEmpty
	}

	// TODO: Either implement duplicate checking or somehow allow duplicates. Latter is better.
	for _, v := range rec {
		if blank.Is(v.Name) {
			return ErrNameBlank
		}

		b.src.Recipes[v.Name] = v
	}

	return nil
}

// SetAttributes reads the JSON-encoded Attributes from the provided Reader and
// adds them to the receiver's Attribute map.
func (b *Builder) SetAttributes(r io.Reader) error {
	attr, err := readAttributes(r)
	if err != nil {
		return errors.Wrap(err, "cannot read Attributes")
	}

	if err := b.setAttributes(attr); err != nil {
		return errors.Wrap(err, "cannot set Attributes")
	}

	return nil
}

// setAttributes adds the provided Attributes to the receiver's Attribute map.
func (b *Builder) setAttributes(attr []Attribute) error {
	if len(attr) <= 0 {
		return ErrSetterEmpty
	}

	for _, v := range attr {
		if blank.Is(v.Name) {
			return ErrNameBlank
		}

		b.src.Attributes[v.Name] = v
	}

	return nil
}

// SetAttrGroups reads the JSON-encoded AttrGroups from the provided
// Reader and adds them to the receiver's AttrGroup map.
func (b *Builder) SetAttrGroups(r io.Reader) error {
	grp, err := readAttrGroups(r)
	if err != nil {
		return errors.Wrap(err, "cannot read AttrGroups")
	}

	if err := b.setAttrGroups(grp); err != nil {
		return errors.Wrap(err, "cannot set AttrGroups")
	}

	return nil
}

// setAttrGroups adds the provided AttrGroups to the receiver's
// AttrGroup map.
func (b *Builder) setAttrGroups(grp []AttrGroup) error {
	if len(grp) <= 0 {
		return ErrSetterEmpty
	}

	for _, v := range grp {
		if blank.Is(v.Name) {
			return ErrNameBlank
		}

		b.src.AttrGroups[v.Name] = v
	}

	return nil
}
