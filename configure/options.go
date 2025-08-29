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

// FuncOption is a generic constraint that permits any function type
// whose underlying type is func(*T). This enables the top-level Apply function
// to accept custom-defined option types, such as `type MyOption func(*T)`.
type FuncOption[T any] interface {
	~func(*T)
}

// FuncOptionE is a generic constraint that permits any function type
// whose underlying type is func(*T) error. This enables the top-level ApplyE
// function to accept custom-defined, error-returning option types.
type FuncOptionE[T any] interface {
	~func(*T) error
}

// AnyOption is a generic constraint that permits any function type
// whose underlying type is either func(*T) or func(*T) error.
// This provides a convenient way to create functions that can accept
// both error-returning and non-error-returning function options.
type AnyOption[T any] interface {
	FuncOptionE[T] | FuncOption[T] | any
}
