package schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// MapMutator is an interface that describes something that can mutate a schema.Schema
type MapMutator interface {
	AddTo(map[string]*schema.Schema) error
}
