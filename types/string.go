/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

// String is an interface that represents a string-like type.
// It can be a string, a byte slice, or a rune slice.
type String interface {
	Object
	~string | ~[]byte | ~[]rune
}

// StringLike converts a string-like type to a string.
func StringLike[T String](t T) string {
	return string(t)
}
