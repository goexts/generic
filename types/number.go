// Package types defines interfaces for different types of numbers.
package types

// Float is an interface that represents a float32 or float64.
type Float interface {
	~float32 | ~float64 // The tilde (~) operator is used to specify the underlying type of a type parameter.
}

// UnsignedInteger is an interface that represents an unsigned integer type.
type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer is an interface that represents a signed integer type.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Complex is an interface that represents a complex64 or complex128 number.
type Complex interface {
	~complex64 | ~complex128
}

// Number is an interface that represents any number type.
// It includes all the interfaces defined above.
type Number interface {
	Float | UnsignedInteger | Integer | Complex
}
