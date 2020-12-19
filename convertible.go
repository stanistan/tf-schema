package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// ConvertResourceMap converts a map of ResourceConveribles into a map of schema.Resources.
func ConvertResourceMap(in map[string]ResourceConverible) map[string]*schema.Resource {
	out := make(map[string]*schema.Resource, len(in))
	for k, v := range in {
		out[k] = v.AsResource()
	}
	return out
}
