package check

// IsZero returns true if the value is zero.
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}
