package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func TestOption_Optional(t *testing.T) {
	assertOption(t, &schema.Schema{Optional: true}, Optional)
}

func TestOption_Required(t *testing.T) {
	assertOption(t, &schema.Schema{Required: true}, Required)
}

func TestOption_ListOf(t *testing.T) {
	expected := &schema.Schema{
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Type: schema.TypeList,
	}
	assertOption(t, expected, ListOf(String))
	assertOption(t, expected, ListOf(&schema.Schema{Type: schema.TypeString}))
}

func TestOption_MapOf(t *testing.T) {
	expected := &schema.Schema{
		Elem: &schema.Schema{Type: schema.TypeString},
		Type: schema.TypeMap,
	}
	assertOption(t, expected, MapOf(String))
	assertOption(t, expected, MapOf(&schema.Schema{Type: schema.TypeString}))
}

func assertOption(t *testing.T, expected *schema.Schema, opt Option, args ...interface{}) {
	t.Helper()
	s := &schema.Schema{}
	opt(s)
	assert.Equal(t, expected, s, args...)
}
