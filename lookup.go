package generic

// MustLookup is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it panics with the error message.
func MustLookup[T any](v T, ok bool) T {
	if !ok {
		panic("key not found")
	}
	return v
}

// MustLookupOr is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it returns the default value.
func MustLookupOr[T any](def T, v T, ok bool) T {
	if !ok {
		return def
	}
	return v
}

// MustLookupOrZero is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it returns a zero value.
func MustLookupOrZero[T any](v T, ok bool) T {
	if !ok {
		return *new(T)
	}
	return v
}

// MustLookupOrNil is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it returns nil.
func MustLookupOrNil[T any](v *T, ok bool) *T {
	if !ok {
		return nil
	}
	return v
}
