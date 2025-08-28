// Package cmp provides utility functions for comparing ordered types.
package cmp

import "golang.org/x/exp/constraints"

// Compare returns -1, 0, or 1 if a is less than, equal to, or greater than b, respectively.
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

// Clamp returns v clamped to the range [lo, hi].
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
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}
