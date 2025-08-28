/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package cast provides utility functions for type casting and assertion.
package cast

import (
	"fmt"
)

// Try attempts to convert a value to the specified type.
// If the conversion is successful, it returns the converted value and true.
// If the conversion fails, it returns the original value and false.
func Try[T any](v any) (ret T, ok bool) {
	if ret, ok = v.(T); ok {
		return ret, true
	}
	return ret, false
}

// Or attempts to convert a value to the specified type.
// If the conversion is successful, it returns the converted value.
// If the conversion fails, it returns the default value.
func Or[T any](v any, def T) T {
	if ret, ok := v.(T); ok {
		return ret
	}
	return def
}

// OrZero attempts to convert a value to the specified type.
// If the conversion is successful, it returns the converted value.
// If the conversion fails, it returns a zero value of the specified type.
func OrZero[T comparable](v any) (zero T) {
	if ret, ok := v.(T); ok {
		return ret
	}
	return zero
}

// Must attempts to convert a value to the specified type.
// If the conversion is successful, it returns the converted value.
// If the conversion fails, it panics with the message.
func Must[T any](v any) T {
	if ret, ok := v.(T); ok {
		return ret
	}
	panic(fmt.Sprintf("value is not type of %T", v))
}
