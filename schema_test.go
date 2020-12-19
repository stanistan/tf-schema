package schema_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func TestNamedSchema(t *testing.T) {

	bar := ComputedString("bar")
	ListOfBarResource := Type(ListOf(Resource(bar)))
	s := Resource(
		OptionalString("foo"),
		RequiredBool("ring"),
		ListOfBarResource("bars", Computed),
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
