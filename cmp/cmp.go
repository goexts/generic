package cmp

import "golang.org/x/exp/constraints"

// Compare returns an integer comparing two values.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
// This function is designed to be fully compatible with the standard library's
// `slices.SortFunc`, making it a convenient tool for sorting slices of any
// ordered type.
func Compare[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// Min returns the smaller of a or b.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of a or b.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Clamp returns v clamped to the inclusive range [lo, hi].
// If v is less than lo, it returns lo.
// If v is greater than hi, it returns hi.
// Otherwise, it returns v.
func Clamp[T constraints.Ordered](v, lo, hi T) T {
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
