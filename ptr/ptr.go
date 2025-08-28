/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ptr provides utility functions for working with pointers.
package ptr

// Of returns a pointer to the given value.
func Of[T any](v T) *T {
	return &v
}

// Val returns the value of the pointer.
// If the pointer is nil, it returns the zero value of the type.
func Val[T any](v *T) T {
	if v != nil {
		return *v
	}
	var zero T
	return zero
}

// To attempts to convert an any value to a pointer of type *T.
// If v is of type T, it returns a pointer to it.
// If v is already *T, it returns v directly.
// Otherwise, it returns a pointer to a new zero value of T.
func To[T any](v any) *T {
	// Check if v is of type T and return its address
	if val, ok := v.(T); ok {
		return &val
	}
	// Check if v is already a *T pointer and return it
	if ptr, ok := v.(*T); ok {
		return ptr
	}
	return new(T)
}

// ToVal attempts to convert an any value to a value of type T.
// If v is a non-nil pointer *T, it returns the dereferenced value.
// If v is of type T, it returns v directly.
// Otherwise, it returns the zero value of T.
func ToVal[T any](v any) T {
	// Handle non-nil pointer case
	if ptr, ok := v.(*T); ok && ptr != nil {
		return *ptr
	}
	// Handle direct value case
	if val, ok := v.(T); ok {
		return val
	}
	// Return zero value as fallback
	return *new(T)
}
