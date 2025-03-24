/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package settings provides type-safe configuration management with default value handling.
package settings

// ApplyFunc is a ApplyFunc function for Apply
// Decrypted: use Func instead of ApplyFunc. Will be removed in v0.3.0
type ApplyFunc[S any] func(*S)

// Apply executes the configuration function with nil safety.
// Decrypted: use Func instead of ApplyFunc. Will be removed in v0.3.0
func (s ApplyFunc[S]) Apply(v *S) {
	if s == nil {
		return
	}
	(s)(v)
}

// ApplySetting defines the interface for configuration settings.
// Decrypted: use Applier instead of ApplySetting. Will be removed in v0.3.0
type ApplySetting[S any] interface {
	Apply(v *S)
}

// Applier defines the interface for configuration application mechanisms.
type Applier[S any] interface {
	// Apply modifies the target struct with specific settings
	Apply(target *S)
}

// ApplierE defines enhanced configuration interface with error handling
type ApplierE[S any] interface {
	Apply(target *S) error
}

// Func defines the standard function type for configuration operations.
type Func[S any] func(*S)

//type FuncA[S any] = func(*S)

// Apply executes the configuration function with nil safety.
func (f Func[S]) Apply(target *S) {
	if f != nil {
		f(target)
	}
}

// FuncE defines enhanced configuration function with error return
type FuncE[S any] func(*S) error

// Apply executes the configuration function with nil safety.
// Returns:
//   - error: Any error encountered during execution
//   - nil: Success or non-error configuration applied
func (f FuncE[S]) Apply(target *S) error {
	if f != nil {
		return f(target)
	}
	return nil
}

// -------------------------
// Unified Application Logic
// -------------------------
func apply[S any](target *S, setting any) bool {
	switch s := setting.(type) {
	case func(*S):
		Func[S](s).Apply(target)
	case Func[S]:
		s.Apply(target)
	case Applier[S]:
		s.Apply(target)
	default:
		return false
	}
	return true
}

func applyWithError[S any](target *S, setting any) (bool, error) {
	var err error
	switch s := setting.(type) {
	case func(*S) error:
		err = FuncE[S](s).Apply(target)
	case FuncE[S]:
		err = s.Apply(target)
	case ApplierE[S]:
		err = s.Apply(target)
	default:
		return false, newConfigError(ErrUnsupportedType, setting, nil)
	}
	return true, wrapIfNeeded(err, setting)
}

func mixedApply[S any](target *S, setting any) error {
	applied, err := applyWithError(target, setting)
	// if not applied will not return. try down apply
	if applied {
		if err == nil {
			return nil
		}
		return err
	}

	// if applied will continue to next
	if apply(target, setting) {
		return nil
	}

	// all tried failed, return the error
	return err
}

// Apply configures a target struct with ordered settings.
// Parameters:
//   - target: Pointer to the struct being configured (non-nil)
//   - settings: Ordered list of configuration functions
//
// Returns:
//   - *S: Configured struct pointer (same as input)
func Apply[S any](target *S, settings []func(*S)) *S {
	if target == nil {
		return nil
	}
	for _, setting := range settings {
		setting(target)
	}
	return target
}

// ApplyOr is an apply settings with defaults
// Decrypted: use WithDefault instead of ApplyOr. Will be removed in v0.3.0
func ApplyOr[S any](s *S, fs ...func(*S)) *S {
	return Apply(s, fs)
}

// ApplyOrZero is an apply settings with defaults
// Decrypted: use WithZero instead of ApplyOrZero. Will be removed in v0.3.0
func ApplyOrZero[S any](fs ...func(*S)) *S {
	var val S
	return Apply(&val, fs)
}

// ApplyE applies a list of settings to a target struct.
// Parameters:
//   - target: Pointer to the struct being configured (non-nil)
//   - settings: Ordered list of configuration functions
//
// Returns:
//   - *S: Configured struct pointer (same as input)
//   - error: Any error encountered during configuration
func ApplyE[S any](target *S, settings []func(target *S) error) (*S, error) {
	if target == nil {
		return nil, newConfigError(ErrEmptyTargetValue, nil, nil)
	}
	var err error
	for _, setting := range settings {
		if err = setting(target); err != nil {
			return nil, err
		}
	}
	return target, nil
}

// ApplyStrict is a version for strict type safety
// Parameters:
//   - target: Pointer to the struct being configured (non-nil)
//   - settings: Ordered list of configuration functions
//
// Returns:
//   - *S: Configured struct pointer (same as input)
//
// Panics:
//   - If target is nil
//   - If any setting is not a supported type
func ApplyStrict[S any](target *S, settings []any) *S {
	if target == nil {
		panic(newConfigError(ErrEmptyTargetValue, nil, nil))
	}
	for _, setting := range settings {
		if !apply(target, setting) {
			panic(newConfigError(ErrUnsupportedType, setting, nil))
		}
	}
	return target
}

// ApplyStrictE applies a list of settings to a target struct.
// Parameters:
//   - target: Pointer to the struct being configured (non-nil)
//   - settings: Ordered list of configuration functions
//
// Returns:
//   - *S: Configured struct pointer (same as input)
func ApplyStrictE[S any](target *S, settings []any) (*S, error) {
	if target == nil {
		return nil, newConfigError(ErrEmptyTargetValue, nil, nil)
	}
	for _, setting := range settings {
		applied, err := applyWithError(target, setting)
		if !applied {
			return nil, err
		}
	}
	return target, nil
}

// ApplyMixed applies a list of settings to a target struct.
// Parameters:
//   - target: Pointer to the struct being configured (non-nil)
//   - settings: Ordered list of configuration functions
//
// Returns:
//   - *S: Configured struct pointer (same as input)
func ApplyMixed[S any](target *S, settings []any) (*S, error) {
	if target == nil {
		return nil, newConfigError(ErrEmptyTargetValue, nil, nil)
	}
	var err error
	for _, setting := range settings {
		if err = mixedApply(target, setting); err != nil {
			return nil, err
		}
	}
	return target, nil
}

// New creates a zero-value instance and applies settings.
// Parameters:
//   - settings: Configuration functions to apply
//
// Retur
// Returns:
//   - *S: New configured instance
func New[S any](settings []func(*S)) *S {
	var zero S
	for _, setting := range settings {
		_ = apply(&zero, setting)
	}
	return &zero
}

// ApplyAny applies a list of settings to a target struct.
// Deprecated: Use ApplyMixed with slice syntax. Will be removed in v0.3.0
//
//	Before: ApplyAny(&obj, []interface{}{f1, f2})
//	After:  ApplyMixed(&obj, []interface{}{f1, f2}) (*obj, error)
func ApplyAny[S any](target *S, settings []any) *S {
	_, _ = ApplyMixed(target, settings)
	return target
}
