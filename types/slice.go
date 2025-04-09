/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

// remove because it is not supported
// type Array[T any] interface{ ~[...]T }

type Slice[T any] interface {
	~[]T
}
