package configure

// Option represents a function that configures an object of type T.
// It is the primary, non-error-returning type for the Functional Options Pattern.
type Option[T any] func(*T)

// Apply implements the Applier[T] interface, allowing an Option[T] to be used
// as a flexible option type with functions like ApplyAny.
func (o Option[T]) Apply(target *T) {
	if o != nil {
		o(target)
	}
}

// OptionE represents a function that configures an object of type T and may
// return an error. The 'E' suffix is a convention for "Error".
type OptionE[T any] func(*T) error

// Apply implements the ApplierE[T] interface, allowing an OptionE[T] to be used
// as a flexible option type with functions like ApplyAny.
func (o OptionE[T]) Apply(target *T) error {
	if o != nil {
		return o(target)
	}
	return nil
}

// OptionConstraint is a generic constraint that permits any function type
// whose underlying type is func(*T). This enables the top-level Apply function
// to accept custom-defined option types, such as `type MyOption func(*T)`.
type OptionConstraint[T any] interface {
	~func(*T)
}

// OptionEConstraint is a generic constraint that permits any function type
// whose underlying type is func(*T) error. This enables the top-level ApplyE
// function to accept custom-defined, error-returning option types.
type OptionEConstraint[T any] interface {
	~func(*T) error
}
