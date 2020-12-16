package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// ResourceConverible is anything that can be converted to a schema.Resource.
type ResourceConverible interface {
	AsResource() *schema.Resource
}

// ProviderConvertible is anything that can be converted into a schema.Provider.
type ProviderConvertible interface {
	AsProvider() *schema.Provider
}

// ConvertResourceMap converts a map of ResourceConveribles into a map of schema.Resources.
func ConvertResourceMap(in map[string]ResourceConverible) map[string]*schema.Resource {
	out := make(map[string]*schema.Resource, len(in))
	for k, v := range in {
		out[k] = v.AsResource()
	}
	return out
}
