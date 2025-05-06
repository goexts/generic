// Package settings implements the functions, types, and interfaces for the module.
package settings

// Defaulter is an interface that provides a method to apply default settings.
// Decrypted: use Apply instead of Defaulter
type Defaulter interface {
	// Defaults applies the default settings to the implementing type.
	Defaults()
}

// ErrorDefaulter is an interface that provides a method to apply default settings and return an error.
// Decrypted: use ApplyE instead of ErrorDefaulter
type ErrorDefaulter interface {
	// Defaults applies the default settings to the implementing type and returns an error.
	Defaults() error
}

// ApplyDefaults applies the given settings and default settings to the provided value.
// Decrypted: use Apply instead of ApplyDefaults
//
// It first applies the given settings using the Apply function, then checks if the value
// implements the Defaulter interface. If it does, it calls the ApplyDefaults method
// to apply the default settings.
func ApplyDefaults[S any](s *S, fs []func(*S)) *S {
	s = Apply(s, fs)
	switch v := any(s).(type) {
	case ErrorDefaulter:
		_ = v.Defaults()
		return s
	case Defaulter:
		v.Defaults()
		return s
	}
	return s
}

// WithDefault applies settings and handles default values.
// Parameters:
//   - target: Struct pointer to configure
//   - settings: Configuration functions to apply
//
// Returns:
//   - *S: Configured struct pointer
func WithDefault[S any](target S, settings ...any) *S {
	for _, setting := range settings {
		_ = apply(&target, setting)
	}
	return &target
}

// WithZero applies settings and handles default values.
// Parameters:
//   - settings: Configuration functions to apply
//
// Returns:
//   - *S: Configured struct pointer
func WithZero[S any](settings ...any) *S {
	var zero S
	for _, setting := range settings {
		_ = apply(&zero, setting)
	}
	return &zero
}

// WithDefaultE applies settings and handles default values.
// Parameters:
//   - target: Struct pointer to configure
//   - settings: Configuration functions to apply
//
// Returns:
//   - *S: Configured struct pointer
//   - error: Error if any occurred during the configuration process
func WithDefaultE[S any](target *S, settings ...any) (*S, error) {
	for _, setting := range settings {
		if _, err := applyWithError(target, setting); err != nil {
			return nil, err
		}
	}
	return target, nil
}

// WithZeroE applies settings and handles default values.
// Parameters:
//   - settings: Configuration functions to apply
//
// Returns:
//   - *S: Configured struct pointer
//   - error: Error if any occurred during the configuration process
func WithZeroE[S any](settings ...any) (*S, error) {
	var zero S
	var err error
	for _, setting := range settings {
		if _, err = applyWithError(&zero, setting); err != nil {
			return nil, err
		}
	}
	return &zero, nil
}

// WithMixed applies settings and handles default values.
// Parameters:
//   - target: Struct pointer to configure
//   - settings: Configuration functions to apply
//
// Returns:
//   - *S: Configured struct pointer
//   - error: Error if any occurred during the configuration process
func WithMixed[S any](target *S, settings ...any) (*S, error) {
	var err error
	var defaults []any
	for _, setting := range settings {
		switch v := setting.(type) {
		case Defaulter:
			defaults = append(defaults, v)
		case ErrorDefaulter:
			defaults = append(defaults, v)
		default:
			if err = mixedApply(target, setting); err != nil {
				return nil, err
			}
		}
	}
	return WithInterfaces(target, defaults...)
}

func WithInterfaces[S any](target *S, settings ...any) (*S, error) {
	var err error
	for _, setting := range settings {
		switch v := setting.(type) {
		case Defaulter:
			v.Defaults()
		case ErrorDefaulter:
			if err = v.Defaults(); err != nil {
				return nil, err
			}
		default:
			return nil, newConfigError(ErrUnsupportedType, setting, nil)
		}
	}
	return target, nil
}

// WithMixedZero applies settings and handles default values.
// Parameters:
//   - settings: Configuration functions to apply
//
// Returns:
//   - *S: Configured struct pointer
//   - error: Error if any occurred during the configuration process
func WithMixedZero[S any](settings ...any) (*S, error) {
	var zero S
	return WithMixed(&zero, settings...)
}

// ApplyErrorDefaults applies the given settings and default settings to the provided value.
// Decrypted: use WithDefaultE instead of ApplyErrorDefaults
//
// It first applies the given settings using the Apply function, then checks if the value
// implements the ErrorDefaulter interface. If it does, it calls the ApplyDefaults method
func ApplyErrorDefaults[S any](s *S, fs []func(*S)) (*S, error) {
	s = Apply(s, fs)
	switch v := any(s).(type) {
	case ErrorDefaulter:
		err := v.Defaults()
		if err != nil {
			return nil, err
		}
		return s, nil
	case Defaulter:
		v.Defaults()
		return s, nil
	}
	return s, nil
}

// ApplyDefaultsOr applies the given settings and default settings to the provided value.
// Decrypted: use WithZeroE instead of ApplyDefaultsOr
//
// It is a convenience wrapper around ApplyDefaults that accepts a variable number of setting functions.
func ApplyDefaultsOr[S any](s *S, fs ...func(*S)) *S {
	// Call ApplyDefaults with the provided value and setting functions
	return ApplyDefaults(s, fs)
}

// ApplyErrorDefaultsOr applies the given settings and default settings to the provided value.
// Decrypted: use WithDefaultE instead of ApplyErrorDefaultsOr
//
// It is a convenience wrapper around ApplyDefaults that accepts a variable number of interface{} values.
func ApplyErrorDefaultsOr[S any](s *S, fs ...func(*S)) (*S, error) {
	return ApplyErrorDefaults(s, fs)
}

// ApplyDefaultsOrZero applies the given settings and default settings to a zero value of the type.
// Decrypted: use WithZero instead of ApplyDefaultsOrZero
//
// It creates a zero value of the type, then calls ApplyDefaults to apply the given settings and default settings.
func ApplyDefaultsOrZero[S any](fs ...func(*S)) *S {
	var zero S
	return ApplyDefaults(&zero, fs)
}

// ApplyErrorDefaultsOrZero applies the given settings and default settings to a zero value of the type.
// Decrypted: use WithZeroE instead of ApplyErrorDefaultsOrZero
//
// It creates a zero value of the type, then calls ApplyDefaults to apply the given settings and default settings.
func ApplyErrorDefaultsOrZero[S any](fs ...func(*S)) (*S, error) {
	var zero S
	return ApplyErrorDefaults(&zero, fs)
}

// ApplyDefaultsOrError applies the given settings and default settings to the provided value.
// Decrypted: use WithDefaultE instead of ApplyDefaultsOrError
//
// It is a convenience wrapper around ApplyDefaults that accepts a variable number of setting functions.
func ApplyDefaultsOrError[S any](s *S, fs ...func(*S)) (*S, error) {
	// Call ApplyDefaults with the provided value and setting functions
	return ApplyErrorDefaults(s, fs)
}
