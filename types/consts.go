/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

// Const is an interface that represents a value that is either a Number, a Boolean, or a string.
// It is used to define constants in Go.
type Const interface {
	Number | Boolean | ~string
}

type Object interface {
	any
}
