package configure

// Applier is an interface for types that can apply a configuration to an object.
// It provides an extension point for ApplyAny, allowing custom types to be
// used as options without reflection.
type Applier[T any] interface {
	Apply(*T)
}

// ApplierE is an interface for types that can apply a configuration and return an error.
// It provides an extension point for ApplyAny, allowing custom types to be
// used as options without reflection.
type ApplierE[T any] interface {
	Apply(*T) error
}

// apply is a private helper that applies a single non-error-returning option.
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
func applyAny[T any](target *T, opt any) error {
	applied, err := applyE(target, opt)
	if applied {
		return err
	}
	if apply(target, opt) {
		return nil
	}
	return err
}

// Apply applies a slice of options to the target.
// It is the core, high-performance function for applying a homogeneous set of
// type-safe options. Its generic constraint allows for custom-defined option
// types, such as `type MyOption func(*T)`.
//
// For handling mixed option types, see ApplyAny.
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
// It is the core, high-performance function for applying a homogeneous set of
// type-safe, error-returning options. Its generic constraint allows for
// custom-defined option types.
//
// For handling mixed option types, see ApplyAny.
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
// This function provides flexibility by using type assertions to handle
// heterogeneous options, at the cost of compile-time type safety and a minor
// performance overhead.
//
// For type-safe, high-performance application, see Apply or ApplyE.
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

// New creates a new instance of T, applies the given options, and returns it.
// It is a convenient top-level constructor for simple object creation where the
// configuration type and the product type are the same.
//
// It uses ApplyAnyWith for maximum flexibility in accepting options.
func New[T any](opts ...any) (*T, error) {
	var zero T
	return ApplyAnyWith(&zero, opts...)
}
