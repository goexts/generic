package cast

// Try attempts to perform a type assertion of v to type T.
// It is a direct, generic wrapper around Go's `value, ok` idiom.
// If the assertion is successful, it returns the converted value and true.
// If the assertion fails, it returns the zero value of T and false.
func Try[T any](v any) (T, bool) {
	val, ok := v.(T)
	return val, ok
}

// Or attempts to perform a type assertion of v to type T.
// If the assertion is successful, it returns the converted value.
// If the assertion fails, it returns the provided default value `def`.
func Or[T any](v any, def T) T {
	if ret, ok := v.(T); ok {
		return ret
	}
	return def
}

// OrZero attempts to perform a type assertion of v to type T.
// If the assertion is successful, it returns the converted value.
// If the assertion fails, it returns the zero value of type T.
func OrZero[T any](v any) T {
	if ret, ok := v.(T); ok {
		return ret
	}
	var zero T
	return zero
}
