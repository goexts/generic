/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

// Pointer returns a pointer to value of Code.
func Pointer[T any](v T) *T {
	return &v
}

// Value returns the value of the pointer of Code.
func Value[T any](v *T) T {
	if v != nil {
		return *v
	}
	var zero T
	return zero
}

// PointerOrZero attempts to convert the input value to type *T.
// Returns a pointer to v if v is of type T, returns v directly if it's already *T,
// otherwise returns a pointer to a zero value of T.
// Parameters:
//
//	v: The value to be converted, of any type.
//
// Returns:
//
//	A pointer of type *T pointing to v or a newly allocated zero value.
func PointerOrZero[T any](v any) *T {
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

// ValueOrZero attempts to convert the input value to the specified generic type.
// Returns the value if:
// - input is a non-nil pointer to the target type (*T)
// - input is a direct value of the target type (T)
// Otherwise returns the zero value of type T.
// Parameters:
//
//	v: any type input value that needs conversion
//
// Returns:
//
//	T value of the target type, or zero value when conversion fails
func ValueOrZero[T any](v any) T {
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
