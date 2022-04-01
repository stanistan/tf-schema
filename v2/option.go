package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

// Optional is an Option that marks the scheam Optional.
func Optional(s *schema.Schema) {
	s.Optional = true
}

// Required is an Option that marks the schema Required.
func Required(s *schema.Schema) {
	s.Required = true
}

// Computed is an Option that marks the schema Computed.
func Computed(s *schema.Schema) {
	s.Computed = true
}

// String is an Option that marks the schema to be a String.
func String(s *schema.Schema) {
	s.Type = schema.TypeString
}

// Bool is an Option that marks the schema to be a Bool.
func Bool(s *schema.Schema) {
	s.Type = schema.TypeBool
}

// Int is an Option that marks the schema to be an Int.
func Int(s *schema.Schema) {
	s.Type = schema.TypeInt
}

// List is an Option that marks the schema to be a List.
func List(s *schema.Schema) {
	s.Type = schema.TypeList
}

// Elem creates an Option that sets the Elem field on the schema.
//
// Accepts either schema.Schema, schema.Response, or an Option.
func Elem(r interface{}) Option {
	return func(s *schema.Schema) {
		s.Elem = toElemType(r)
	}
}

// Default creates an Option marks the schema's Default field.
func Default(v interface{}) Option {
	return func(s *schema.Schema) {
		s.Default = v
	}
}

// ListOf is an Option that marks the schema to be a List of value.
func ListOf(value interface{}) Option {
	return Options(List, Elem(value))
}

// Map is an Option that marks the schema to be a Map.
func Map(s *schema.Schema) {
	s.Type = schema.TypeMap
}

// MapOf is an Option that marks the schema to be a map of value.
func MapOf(value interface{}) Option {
	return Options(Map, Elem(value))
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
		return v // this will fail in TF itself
	}
}
