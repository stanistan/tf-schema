package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

// Schema is a wrapper for the terraform map[string]*schema.Schema.
//
// This is similar to something you would use if you were creating
// a resource.
type Schema struct {
	s map[string]*schema.Schema
}

// AsResource turns this schema into a schema.Resource.
func (s *Schema) AsResource() *schema.Resource {
	return &schema.Resource{
		Schema: s.Get(),
	}
}

// New creates a Schema given any MapMutators and applies them
func New(fs ...MapMutator) (*Schema, error) {
	schema := &Schema{map[string]*schema.Schema{}}
	return schema.Push(fs...)
}

// Must creates a Schema and panics if there is an error.
func Must(fs ...MapMutator) *Schema {
	s, err := New(fs...)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return s
}

// Resource is an alias for Must
var Resource = Must

// Push adds any MapMutators to a Schema.
func (s *Schema) Push(fs ...MapMutator) (*Schema, error) {
	for _, f := range fs {
		if err := f.AddTo(s.s); err != nil {
			return s, errors.WithStack(err)
		}
	}
	return s, nil
}

// MustPush adds any MapMutators to a Schema, panics on failure.
func (s *Schema) MustPush(fs ...MapMutator) *Schema {
	_, err := s.Push(fs...)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return s
}

// Get extracts the schema map from Schema.
func (s *Schema) Get() map[string]*schema.Schema {
	return s.s
}

// ComputedList wraps this Schema into a MapMutator.
func (s *Schema) ComputedList(name string) MapMutator {
	return NewNamed(name, List, Computed, Elem(s.AsResource()))
}

// ComputedMap wraps this Schema into a MapMutator.
func (s *Schema) ComputedMap(name string) MapMutator {
	return NewNamed(name, Map, Computed, Elem(s.AsResource()))
}
