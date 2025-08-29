package configure

import (
	"reflect"
)

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

// applyReflect uses reflection to apply a function-based option.
// It returns whether the option was applied and any error that occurred.
func applyReflect[T any](target *T, opt any) (bool, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() != reflect.Func {
		return false, newConfigError(ErrUnsupportedType, opt, nil)
	}

	// Check for error-returning function: func(*T) error
	if v.Type().ConvertibleTo(reflect.TypeOf(OptionE[T](nil))) {
		converted := v.Convert(reflect.TypeOf(OptionE[T](nil))).Interface().(OptionE[T])
		err := converted(target)
		if err != nil {
			return true, newConfigError(ErrExecutionFailed, opt, err)
		}
		return true, nil
	}

	// Check for non-error-returning function: func(*T)
	if v.Type().ConvertibleTo(reflect.TypeOf(Option[T](nil))) {
		converted := v.Convert(reflect.TypeOf(Option[T](nil))).Interface().(Option[T])
		converted(target)
		return true, nil
	}

	return false, nil
}

// applyAny is a private helper that attempts to apply an option of unknown type.
func applyAny[T any](target *T, opt any) error {
	// 1. Try ApplierE (error-returning, type-asserted)
	applied, err := applyE(target, opt)
	if applied {
		return err
	}

	// 2. Try Applier (non-error-returning, type-asserted)
	if apply(target, opt) {
		return nil
	}

	// 3. Fallback to reflection for convertible function types
	applied, err = applyReflect(target, opt)
	if applied {
		return err
	}

	// If nothing applied, return the original error from applyE,
	// which indicates an unsupported type.
	return err
}

// Apply applies a slice of options to the target.
// It is the core, high-performance function for applying a homogeneous set of
// type-safe options. Its generic constraint allows for custom-defined option
// types, such as `type MyOption func(*T)`.
//
// For handling mixed option types, see ApplyAny.
func Apply[T any, O FuncOption[T]](target *T, opts []O) *T {
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
func ApplyE[T any, O FuncOptionE[T]](target *T, opts []O) (*T, error) {
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
// This function provides flexibility by using reflection to handle
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

// OptionSet bundles multiple options into a single option.
// This allows for creating reusable and modular sets of configurations.
func OptionSet[T any](opts ...Option[T]) Option[T] {
	return func(t *T) {
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

// New creates a new instance of T, applies the given options, and returns it.
// It is a convenient top-level constructor for simple object creation where the
// configuration type and the product type are the same.
//
// It uses ApplyAnyWith for maximum flexibility in accepting options.
func New[T any](opts ...any) (*T, error) {
	var zero T
	return ApplyAny(&zero, opts)
}
