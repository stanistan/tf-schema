package schema

var (
	// Required schemas
	RequiredBool   = Type(Required, Bool)
	RequiredInt    = Type(Required, Int)
	RequiredString = Type(Required, String)

	// Computed schemas
	ComputedBool   = Type(Computed, Bool)
	ComputedInt    = Type(Computed, Int)
	ComputedString = Type(Computed, String)

	// Optional schemas
	OptionalBool   = Type(Optional, Bool)
	OptionalInt    = Type(Optional, Int)
	OptionalString = Type(Optional, String)
)
