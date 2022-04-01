package schema_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	s "github.com/stanistan/tf-schema/v2"
)

func TestOption_Optional(t *testing.T) {
	assertOption(t, &schema.Schema{Optional: true}, s.Optional)
}

func TestOption_Required(t *testing.T) {
	assertOption(t, &schema.Schema{Required: true}, s.Required)
}

func TestOption_ListOf(t *testing.T) {
	expected := &schema.Schema{
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Type: schema.TypeList,
	}
	assertOption(t, expected, s.ListOf(s.String))
	assertOption(t, expected, s.ListOf(&schema.Schema{Type: schema.TypeString}))
}

func TestOption_MapOf(t *testing.T) {
	expected := &schema.Schema{
		Elem: &schema.Schema{Type: schema.TypeString},
		Type: schema.TypeMap,
	}
	assertOption(t, expected, s.MapOf(s.String), "MapOf extracts correctly")
	assertOption(t, expected, s.MapOf(&schema.Schema{Type: schema.TypeString}), "MapOf supports the raw helper/schema")
}

func assertOption(t *testing.T, expected *schema.Schema, opt s.Option, args ...interface{}) {
	t.Helper()
	s := &schema.Schema{}
	opt(s)
	assert.Equal(t, expected, s, args...)
}
