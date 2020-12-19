package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// MapMutator is an interface that describes something that can mutate a schema.Schema.
type MapMutator interface {
	AddTo(map[string]*schema.Schema) error
}

// ResourceConverible is anything that can be converted to a schema.Resource.
type ResourceConverible interface {
	AsResource() *schema.Resource
}

// ProviderConvertible is anything that can be converted into a schema.Provider.
type ProviderConvertible interface {
	AsProvider() *schema.Provider
}
