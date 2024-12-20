/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package trans implements the functions, types, and interfaces for the module.
package trans

// Pointer returns a pointer to value of Type.
func Pointer[T any](v T) *T {
	return &v
}

// Value returns the value of the pointer of Type.
func Value[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}
