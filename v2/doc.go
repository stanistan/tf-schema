// tf-schema is a library to help trim down the verbosity of writing terraform providers.
//
// It does this in a couple of ways.
//
// 1. By providing primitive combinators to make writing schemas/resources simpler,
// easier to read and write.
//
// 2. By providing a codegen tool so a provider can have its own typed resources,
// not relying on interface{}, but whatever userland type is desired.
//
// These two features can be used separately or together.
//
// See the README at https://github.com/stanistan/tf-schema for examples.
package schema
