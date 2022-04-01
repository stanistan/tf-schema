# tf-schema

A libary to help trim down the verbosity of writing TF providers.

- Generates a typed `Resource` and `Provider` for your plugin
- Provides combinators for declarative and terse schema

If you want to use v2 Providers, you can use `tf-schema/v2`!

## Usage: Codegen & Setup

You are generally writing a provider to interact with some client library.
`tf-schema` works with `go generate` to build typed versions of resources
and providers that are interchangeable with standard ones from terraform.

### Example

#### imports

```go
import (
	lib "some/url/package/lib"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	terraform "github.com/hashicorp/terraform-plugin-sdk/terraform"
	schema "github.com/stanistan/tf-schema"
)
```

#### `go generate`

_This is mandatory to generate the types to interact with the client._

```go
//go:generate go run -mod readonly github.com/stanistan/tf-schema/v2/cmd/generate lib.Client some/url/package/lib
```

#### Provider

```go
func Provider() *tfschema.Provider {
	// libProvider is a generated struct private to your code
	return (&libProvider{
		Schema: schema.Must(
		// provider settings
		),
		// NOTE: the below return interface objects, not just structs
		DataSourcesMap: map[string]schema.ResourceConvertible{
			"lib_foo": dataSourceLibFoo(),
		},
		// NOTE: the below is now _typed_! to the client/struct type you want to be
		// returning, and _not_ just returning an `interface{}`.
		ConfigureFunc: func(d *tfschema.ResourceData) (*lib.Client, error) {
			return lib.NewClient()
		},
	}).AsProvider()
}
```

#### A Datasource/Resource

```go
func dataSourceLibFoo() schema.ResourceConvertible {
	// libResource is a generated struct private to your code
	return &libResource{
		Schema: schema.Must(
			schema.RequiredString("key"),
			schema.OptionalInt("query_page"),
			schema.ComputedString("value"),
		),
		// NOTE: Read is now _typed_ to the _kind_ of `meta` you would be returning
		// from the ConfigureFunc in the provider, and not an interface you have to cast
		// yourself. The validation has already been done.
		Read: func(d *tfschema.ResourceData, c *lib.Client) error {
			// snip...
			data, err := c.GetFoo(d.Get("key").(string))
			// snip...
			return nil
		},
	}
}
```

## Usage: schema Combinators

### `Option`, `Named`, and `Type`

An `Option` is anything that can mutate a `schema.Schema` (tf).

```go
type Option func(*schema.Schema)
```

You can create your own types by using the defining `Option`s, and common
ones are defined in [`option.go`][options] and common types are defined
in [`types.go`][types] using `Type`.

For example:

```go
// import (
//	schema "github.com/stanistan/tf-schema/v2"
//)

// This is how schema.OptionalString is defined.
//
// It returns a schema.NamedSchemaFactory,
// to which you can provide a name, and other schema.Options.
var optionalString = schema.Type(schema.Optional, schema.String)

var optionalStringDefaultsToBanana = schema.Type(
	schema.Optional, schema.String, schema.Default("banana"))

// and later use this with schema.Schema
schema.Must(
	optionalStringDefaultsToBanana("fruit"),
)

// this is equivalent
schema.Must(
    schema.OptionalString("fruit", schema.Default("banana")),
)
```

### `MapMutator`

You can create a schema with anything that conforms to `MapMutator`, and
use them in composition.

```go
type Once struct {
    named Named
}

func (s *Once) AddTo(m map[string]*schema.Schema) error {
    name := s.named.Name()
    _, exists := m[name]
    if exists {
        return fmt.Errorf("can't override name %s", name)
    }
    return s.named.AddTo(m)
}
```

or

```go
type Fruit string

func (s *Fruit) AddTo(m map[string]*schema.Schema) error {
    m["fruit"] = &schema.Schema{ /** stuff */ }
    // this can set multiple keys
    m["not_fruit"] = &schema.Schema{ /** stuff */ }
    return nil
}
```

[options]: ./options.go
[types]: ./types.go
