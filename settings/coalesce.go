// Package settings implements the functions, types, and interfaces for the module.
package settings

import (
	"math"

	"github.com/goexts/generic/types"
)

const epsilon = 1e-9

func Coalesce[T comparable](a, b T) T {
	var zero T
	if a != zero {
		return a
	}
	return b
}

func CoalesceString[T types.String](a, b T) T {
	return T(Coalesce(string(a), string(b)))
}

func CoalesceInt[T types.Integer](a, b T) T {
	return Coalesce(a, b)
}

func CoalesceFloat[T types.Float](a, b T) T {
	if math.IsNaN(float64(a)) || math.Abs(float64(a)) < epsilon {
		return b
	}
	return a
}

func CoalesceBool[T types.Boolean](a, b T) T {
	return Coalesce(a, b)
}
