package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Option can mutate the schema.
//
// Use this for composition.
type Option func(*schema.Schema)

// Options returns an Option that applies the given options.
func Options(opts ...Option) Option {
	return func(s *schema.Schema) {
		for _, opt := range opts {
			opt(s)
		}
	}
}

// Optional marks the scheam Optional
func Optional(s *schema.Schema) {
	s.Optional = true
}

// Required marks the schema Required
func Required(s *schema.Schema) {
	s.Required = true
}

// Computed marks the schema Computed
func Computed(s *schema.Schema) {
	s.Computed = true
}

// String marks the schema to be a String
func String(s *schema.Schema) {
	s.Type = schema.TypeString
}

// Bool marks the schema to be a Bool
func Bool(s *schema.Schema) {
	s.Type = schema.TypeBool
}

// Int marks the schema to be an Int
func Int(s *schema.Schema) {
	s.Type = schema.TypeInt
}

// List marks the schema to be a List
func List(s *schema.Schema) {
	s.Type = schema.TypeList
}

// ListOf marks the schema to be a List of r.
func ListOf(r interface{}) Option {
	return Options(List, Elem(r))
}

// Map marks the schema to be a Map
func Map(s *schema.Schema) {
	s.Type = schema.TypeMap
}

// MapOf marks the schema to be a map of r.
func MapOf(r interface{}) Option {
	return Options(Map, Elem(r))
}

// Elem creates an Option that sets the Elem field on the schema.
//
// Accepts either schema.Schema, schema.Response, or an Option.
func Elem(r interface{}) Option {
	return func(s *schema.Schema) {
		s.Elem = toElemType(r)
	}
}

// Default creates an Option marks the schema's Default field
func Default(v interface{}) Option {
	return func(s *schema.Schema) {
		s.Default = v
	}
}

func toElemType(val interface{}) interface{} {
	if val == nil {
		return val
	}
	switch v := val.(type) {
	case *schema.Resource:
		return v
	case *schema.Schema:
		return v
	case *Schema:
		return v.AsResource()
	case func(*schema.Schema):
		var elemSchema schema.Schema
		v(&elemSchema)
		return &elemSchema
	default:
		// this will fail in TF itself
		return v
	}
}
