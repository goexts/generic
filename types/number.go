package types

import "github.com/goexts/generic/object"

// Number is an interface for all numeric types.
// It is the equivalent of java.lang.Number.
type Number interface {
	object.Object

	// IntValue returns the value of the number as an int.
	IntValue() int
	// Int64Value returns the value of the number as an int64.
	Int64Value() int64
	// Float32Value returns the value of the number as a float32.
	Float32Value() float32
	// Float64Value returns the value of the number as a float64.
	Float64Value() float64
}
