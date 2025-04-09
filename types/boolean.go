/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

// Boolean is an interface that represents a boolean type.
// It is implemented by the built-in bool type.
type Boolean interface {
	Object
	// ~bool is a type constraint that specifies that the type must be a boolean.
	~bool
}
