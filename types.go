package schema

// These are pre-defined type constructors for Named.
//
// Use them like so:
//
//	Resource(
// 		RequiredBool("bool_field_name")
//	)
//
var (
	RequiredBool   NamedSchemaFactory = Type(Required, Bool)
	RequiredInt                       = Type(Required, Int)
	RequiredString                    = Type(Required, String)
	ComputedBool                      = Type(Computed, Bool)
	ComputedInt                       = Type(Computed, Int)
	ComputedString                    = Type(Computed, String)
	OptionalBool                      = Type(Optional, Bool)
	OptionalInt                       = Type(Optional, Int)
	OptionalString                    = Type(Optional, String)
)
