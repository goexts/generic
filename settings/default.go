// Package settings implements the functions, types, and interfaces for the module.
package settings

// Defaulter is an interface that provides a method to apply default settings.
type Defaulter interface {
	// ApplyDefaults applies the default settings to the implementing type.
	ApplyDefaults()
}

// ApplyDefaults applies the given settings and default settings to the provided value.
//
// It first applies the given settings using the Apply function, then checks if the value
// implements the Defaulter interface. If it does, it calls the ApplyDefaults method
// to apply the default settings.
func ApplyDefaults[S any](s *S, fs []func(*S)) *S {
	s = Apply(s, fs)
	if d, ok := any(s).(Defaulter); ok {
		d.ApplyDefaults()
	}
	return s
}

// ApplyDefaultsOr applies the given settings and default settings to the provided value.
//
// It is a convenience wrapper around ApplyDefaults that accepts a variable number of setting functions.
func ApplyDefaultsOr[S any](s *S, fs ...func(*S)) *S {
	// Call ApplyDefaults with the provided value and setting functions
	return ApplyDefaults(s, fs)
}

// ApplyDefaultsOrZero applies the given settings and default settings to a zero value of the type.
//
// It creates a zero value of the type, then calls ApplyDefaults to apply the given settings and default settings.
func ApplyDefaultsOrZero[S any](fs ...func(*S)) *S {
	var zero S
	return ApplyDefaults(&zero, fs)
}
