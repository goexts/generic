package cmp

import (
	"cmp"
)

// IsZero returns true if the value is zero.
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

// ZeroOr returns def if v is the zero value.
func ZeroOr[T comparable](v T, def T) T {
	var zero T
	if v == zero {
		return def
	}
	return v
}

// Or returns the first non-zero value.
func Or[T comparable](vals ...T) T {
	return cmp.Or(vals...)
}
