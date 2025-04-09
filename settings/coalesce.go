// Package settings provides utility functions for coalescing values of various types.
package settings

import (
	"math"

	"github.com/goexts/generic/types"
)

const epsilon = 1e-9 // Small value for floating point comparisons

// Coalesce returns the first non-zero value of comparable type T
// If a is non-zero, returns a; otherwise returns b
func Coalesce[T comparable](a, b T) T {
	var zero T
	if a != zero {
		return a
	}
	return b
}

// CoalesceString handles string coalescing with type safety
// Converts parameters to string type before comparison
func CoalesceString[T types.String](a, b T) T {
	return T(Coalesce(string(a), string(b)))
}

// CoalesceInt provides type-safe coalescing for integer types
// Returns the first non-zero integer value
func CoalesceInt[T types.Integer](a, b T) T {
	return Coalesce(a, b)
}

// CoalesceFloat handles floating point coalescing with epsilon comparison
// Returns b if a is NaN or its absolute value is below epsilon threshold
func CoalesceFloat[T types.Float](a, b T) T {
	if math.IsNaN(float64(a)) || math.Abs(float64(a)) < epsilon {
		return b
	}
	return a
}

// CoalesceBool provides type-safe coalescing for boolean types
// Returns the first non-zero boolean value
func CoalesceBool[T types.Boolean](a, b T) T {
	return Coalesce(a, b)
}
