/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

import (
	"golang.org/x/exp/constraints"
)

// Float is an interface that represents a float32 or float64.
type Float = constraints.Float

// Unsigned is an interface that represents an unsigned integer type.
type Unsigned = constraints.Unsigned

// Signed is an interface that represents a signed integer type.
type Signed = constraints.Signed

// Integer is an interface that represents an integer type.
type Integer = constraints.Integer

// Complex is an interface that represents a complex64 or complex128 number.
type Complex = constraints.Complex

// Ordered is an interface that represents any ordered type.
type Ordered = constraints.Ordered

// Number is an interface that represents any number type.
// It includes all the interfaces defined above.
type Number interface {
	Float | Integer | Complex
}
