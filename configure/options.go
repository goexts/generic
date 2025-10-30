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

// OptionFunc is a generic constraint that permits any function type
// whose underlying type is func(*T). This enables the top-level Apply function
// to accept custom-defined option types, such as `type MyOption func(*T)`.
type OptionFunc[T any] interface {
	~func(*T)
}

// OptionFuncE is a generic constraint that permits any function type
// whose underlying type is func(*T) error. This enables the top-level ApplyE
// function to accept custom-defined, error-returning option types.
type OptionFuncE[T any] interface {
	~func(*T) error
}

// OptionFuncAny is a generic constraint that permits any function type
// whose underlying type is either func(*T) or func(*T) error.
// This provides a convenient way to create functions that can accept
// both error-returning and non-error-returning function options.
type OptionFuncAny[T any] interface {
	OptionFuncE[T] | OptionFunc[T] | any
}

// OptionSet bundles multiple options into a single option.
// This allows for creating reusable and modular sets of configurations.
func OptionSet[T any](opts ...Option[T]) Option[T] {
	return func(t *T) {
		Apply(t, opts)
	}
}

func Chain[S any, T OptionFunc[S]](opts ...T) T {
	return func(t *S) {
		Apply(t, opts)
	}
}

// OptionSetE bundles multiple error-returning options into a single option.
// If any option in the set returns an error, the application stops and the error is returned.
func OptionSetE[T any](opts ...OptionE[T]) OptionE[T] {
	return func(t *T) error {
		_, err := ApplyE(t, opts)
		return err
	}
}

func ChainE[S any, T OptionFuncE[S]](opts ...T) T {
	return func(t *S) error {
		_, err := ApplyE(t, opts)
		return err
	}
}
