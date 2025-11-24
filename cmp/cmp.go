package cmp

//go:generate adptool .
//go:adapter:package cmp

// Min returns the smaller of a or b.
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of a or b.
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Clamp returns v clamped to the inclusive range [lo, hi].
// If v is less than lo, it returns lo.
// If v is greater than hi, it returns hi.
// Otherwise, it returns v.
func Clamp[T Ordered](v, lo, hi T) T {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// IsZero returns true if v is the zero value for its type.
// It is a generic-safe way to check for zero values.
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}
