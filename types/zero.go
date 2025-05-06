/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

// Zero is the zero value for a type.
func Zero[T any]() (zero T) {

	return zero
}

// ZeroOr returns def if v is the zero value.
// Decrypted: use cmp.ZeroOr instead.
func ZeroOr[T comparable](v T, def T) T {
	if v == Zero[T]() {
		return def
	}
	return v
}
