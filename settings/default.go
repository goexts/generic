// Package settings implements the functions, types, and interfaces for the module.
package settings

type Defaulter interface {
	ApplyDefaults()
}

func ApplyDefaults[S any](s *S, fs []func(*S)) *S {
	s = Apply(s, fs)
	if d, ok := any(s).(Defaulter); ok {
		d.ApplyDefaults()
	}
	return s
}

func ApplyDefaultsOrZero[S any](fs ...func(*S)) *S {
	var zero S
	return ApplyDefaults(&zero, fs)
}
