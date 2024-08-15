package generic

// Must is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it panics with the error message.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Must2 is a utility function that ensures a value is not nil and returns it.
func Must2[T any, U any](v T, u U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return v, u
}

// MustOr is a utility function that ensures a value is not nil and returns it.
// If the error is not nil, it returns the default value.
func MustOr[T any](def T, v T, err error) T {
	if err != nil {
		return def
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

// OrNil is a utility function that ensures a value is not nil and returns it.
func OrNil[T any](_ T, err error) error {
	if err != nil {
		return err
	}
	return nil
}

// OrNil2 is a utility function that ensures a value is not nil and returns it.
func OrNil2[T any](_, _ T, err error) error {
	if err != nil {
		return err
	}
	return nil
}
