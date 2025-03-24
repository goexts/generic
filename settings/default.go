// Package settings implements the functions, types, and interfaces for the module.
package settings

// Defaulter is an interface that provides a method to apply default settings.
// Decrypted: use Apply instead of Defaulter
type Defaulter interface {
	// ApplyDefaults applies the default settings to the implementing type.
	ApplyDefaults()
}

// ErrorDefaulter is an interface that provides a method to apply default settings and return an error.
// Decrypted: use ApplyE instead of ErrorDefaulter
type ErrorDefaulter interface {
	// ApplyDefaults applies the default settings to the implementing type and returns an error.
	ApplyDefaults() error
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
		_ = v.ApplyDefaults()
		return s
	case Defaulter:
		v.ApplyDefaults()
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
func WithDefault[S any](target *S, settings ...any) *S {
	for _, setting := range settings {
		_ = apply(target, setting)
	}
	return target
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

func WithDefaultE[S any](target *S, settings ...any) (*S, error) {
	for _, setting := range settings {
		if _, err := applyWithError(target, setting); err != nil {
			return nil, err
		}
	}
	return target, nil
}

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

func MixedDefaultE[S any](target *S, settings ...any) (*S, error) {
	for _, setting := range settings {
		applied, err := applyWithError(target, setting)
		// if not applied will not return. try down apply
		if applied {
			if err == nil {
				continue
			}
			return nil, err
		}

		// if applied will continue to next
		if apply(target, setting) {
			continue
		}

		// all tried failed, return the error
		return nil, err
	}
	return target, nil
}

func MixedZeroE[S any](settings ...any) (*S, error) {
	var zero S
	for _, setting := range settings {
		applied, err := applyWithError(&zero, setting)
		// if not applied will not return. try down apply
		if applied {
			if err == nil {
				continue
			}
			return nil, err
		}

		// if applied will continue to next
		if apply(&zero, setting) {
			continue
		}

		// all tried failed, return the error
		return nil, err
	}
	return &zero, nil
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
		err := v.ApplyDefaults()
		if err != nil {
			return nil, err
		}
		return s, nil
	case Defaulter:
		v.ApplyDefaults()
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
