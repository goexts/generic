package generic

// Must is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it panics with the error message.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// MustOrZero is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it returns a zero value.
func MustOrZero[T any](v T, err error) T {
	if err != nil {
		return *new(T)
	}
	return v
}

// MustOrNil is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it returns nil.
func MustOrNil[T any](v *T, err error) *T {
	if err != nil {
		return nil
	}
	return v
}

// MustOrDefault is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it returns the default value.
func MustOrDefault[T any](v T, err error, def T) T {
	if err != nil {
		return def
	}
	return v
}
