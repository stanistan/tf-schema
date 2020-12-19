package schema_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	s "github.com/stanistan/tf-schema"
)

func TestNamedSchema(t *testing.T) {

	bar := s.ComputedString("bar")
	ListOfBarResource := s.Type(s.ListOf(s.Resource(bar)))
	s := s.Resource(
		s.OptionalString("foo"),
		s.RequiredBool("ring"),
		ListOfBarResource("bars", s.Computed),
	)

	assert.Equal(t, map[string]*schema.Schema{
		"foo": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ring": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"bars": {
			Computed: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"bar": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
	}, s.Get())
}
