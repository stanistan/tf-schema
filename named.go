package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// NamedSchemaFactory is a function that creates a Named
type NamedSchemaFactory func(string, ...Option) *Named

// Named holds a schema.Schema with it's name.
type Named struct {
	name   string
	schema *schema.Schema
}

// NewNamed creates a Named schema given any Options
func NewNamed(name string, opts ...Option) *Named {
	return (&Named{name, &schema.Schema{}}).Apply(opts...)
}

// Type is a constructor of NamedSchemaFactory given default Options
func Type(defaultOpts ...Option) NamedSchemaFactory {
	return func(name string, opts ...Option) *Named {
		return NewNamed(name, defaultOpts...).Apply(opts...)
	}
}

// Apply applies any Options to the contained schema.Schema
func (s *Named) Apply(opts ...Option) *Named {
	Options(opts...)(s.schema)
	return s
}

// Name gets us the name of the Named schema.
func (s *Named) Name() string {
	return s.name
}

// AddTo implements the MapMutator interface
func (s *Named) AddTo(m map[string]*schema.Schema) error {
	m[s.name] = s.schema
	return nil
}
