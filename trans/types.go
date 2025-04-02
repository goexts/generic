/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package trans implements the functions, types, and interfaces for the module.
package trans

import (
	"fmt"
)

// Cast attempts to convert a value to the specified type.
// If the conversion is successful, it returns the converted value and true.
// If the conversion fails, it returns the original value and false.
func Cast[T any](v any) (ret T, ok bool) {
	if ret, ok = v.(T); ok {
		return ret, true
	}
	return ret, false
}

// CastOr attempts to convert a value to the specified type.
// If the conversion is successful, it returns the converted value.
// If the conversion fails, it returns the default value.
func CastOr[T any](v any, def T) T {
	if ret, ok := v.(T); ok {
		return ret
	}
	return def
}

// CastOrZero attempts to convert a value to the specified type.
// If the conversion is successful, it returns the converted value.
// If the conversion fails, it returns a zero value of the specified type.
func CastOrZero[T comparable](v any) (zero T) {
	if ret, ok := v.(T); ok {
		return ret
	}
	return zero
}

// MustCast attempts to convert a value to the specified type.
// If the conversion is successful, it returns the converted value.
// If the conversion fails, it panics with the message.
func MustCast[T any](v any) T {
	if ret, ok := v.(T); ok {
		return ret
	}
	panic(fmt.Sprintf("value is not type of %T", v))
}
