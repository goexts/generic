// Package configure provides a robust and type-safe way to apply functional options to objects.
package configure

// Applier is an interface for types that can apply a configuration to an object.
// It provides a way for ApplyAny to handle custom configuration types without reflection.
type Applier[T any] interface {
	Apply(*T)
}

// ApplierE is an interface for types that can apply a configuration and return an error.
// It provides a way for ApplyAny to handle custom configuration types without reflection.
type ApplierE[T any] interface {
	Apply(*T) error
}

// apply is a private helper that applies a single non-error-returning option.
// It uses a type switch to handle various supported option formats.
func apply[T any](target *T, opt any) bool {
	var applier Applier[T]
	switch o := opt.(type) {
	case func(*T):
		applier = Option[T](o)
	case Option[T]:
		applier = o
	case Applier[T]:
		applier = o
	default:
		return false
	}
	applier.Apply(target)
	return true
}

// applyE is a private helper that applies a single error-returning option.
// It uses a type switch to handle various supported option formats.
func applyE[T any](target *T, opt any) (bool, error) {
	var applier ApplierE[T]
	switch o := opt.(type) {
	case func(*T) error:
		applier = OptionE[T](o)
	case OptionE[T]:
		applier = o
	case ApplierE[T]:
		applier = o
	default:
		return false, newConfigError(ErrUnsupportedType, opt, nil)
	}
	err := applier.Apply(target)
	if err != nil {
		return true, newConfigError(ErrExecutionFailed, opt, err)
	}
	return true, nil
}

// applyAny is a private helper that attempts to apply an option of unknown type.
// It first tries error-returning formats, then falls back to non-error formats.
func applyAny[T any](target *T, opt any) error {
	applied, err := applyE(target, opt)
	if applied {
		return err
	}
	if apply(target, opt) {
		return nil
	}
	// If we reach here, 'err' is the ErrUnsupportedType from applyE.
	return err
}

// Apply applies a slice of options to the target.
// This is the core, high-performance function for type-safe options.
func Apply[T any, O OptionConstraint[T]](target *T, opts []O) *T {
	if target == nil {
		return nil
	}
	for _, f := range opts {
		f(target)
	}
	return target
}

// ApplyWith is the variadic convenience wrapper for Apply.
func ApplyWith[T any](target *T, opts ...Option[T]) *T {
	return Apply(target, opts)
}

// ApplyE applies a slice of error-returning options to the target.
// This is the core, high-performance function for type-safe, error-returning options.
func ApplyE[T any, O OptionEConstraint[T]](target *T, opts []O) (*T, error) {
	if target == nil {
		return nil, newConfigError(ErrEmptyTargetValue, nil, nil)
	}
	for _, f := range opts {
		if err := f(target); err != nil {
			return nil, newConfigError(ErrExecutionFailed, f, err)
		}
	}
	return target, nil
}

// ApplyWithE is the variadic convenience wrapper for ApplyE.
func ApplyWithE[T any](target *T, opts ...OptionE[T]) (*T, error) {
	return ApplyE(target, opts)
}

// ApplyAny applies a slice of options of various types (any).
// This is the flexible, reflection-based function for heterogeneous options.
func ApplyAny[T any](target *T, opts []any) (*T, error) {
	if target == nil {
		return nil, newConfigError(ErrEmptyTargetValue, nil, nil)
	}
	for _, opt := range opts {
		if err := applyAny(target, opt); err != nil {
			return nil, err
		}
	}
	return target, nil
}

// ApplyAnyWith is the variadic convenience wrapper for ApplyAny.
func ApplyAnyWith[T any](target *T, opts ...any) (*T, error) {
	return ApplyAny(target, opts)
}

// FromConfig creates a "super option" from a factory function and a set of configuration options.
// It allows using a configuration object `C` to produce a final product `P`,
// while still being compatible with the `New` function.
func FromConfig[C any, P any](factory func(c *C) (*P, error), opts ...any) OptionE[P] {
	return func(p *P) error {
		var cfg C
		if _, err := ApplyAny(&cfg, opts); err != nil {
			return err
		}

		newProduct, err := factory(&cfg)
		if err != nil {
			return err
		}

		// Replace the target product with the newly created one.
		*p = *newProduct
		return nil
	}
}

// New creates a new instance of T, applies the given options, and returns it.
// It uses ApplyAnyWith for maximum flexibility.
func New[T any](opts ...any) (*T, error) {
	var zero T
	return ApplyAnyWith(&zero, opts...)
}
