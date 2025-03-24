/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package settings implements the functions, types, and interfaces for the module.
package settings

// Result defines the unified return type for configuration operations
type Result[S any] interface {
	// Value returns the configured struct pointer
	Value() *S
	// Err returns any error occurred during configuration
	Err() error
	// Unwrap returns both value and error for direct handling
	Unwrap() (*S, error)
	// MustUnwrap panics if error exists
	MustUnwrap() *S
}

// Setting represents a union type for configuration options.
type Setting[S any] interface {
	func(*S) | func(*S) error | Func[S] | FuncE[S]
}

type baseResult[S any] struct {
	val *S
	err error
}

func (r *baseResult[S]) Value() *S           { return r.val }
func (r *baseResult[S]) Err() error          { return r.err }
func (r *baseResult[S]) Unwrap() (*S, error) { return r.val, r.err }
func (r *baseResult[S]) MustUnwrap() *S {
	if r.err != nil {
		panic(r.err)
	}
	return r.val
}

func NewResult[S any](val *S, err error) Result[S] {
	return &baseResult[S]{
		val: val,
		err: err,
	}
}

func NewValueResult[S any](val *S) Result[S] {
	return &baseResult[S]{
		val: val,
	}
}
func NewErrorResult[S any](err error) Result[S] {
	return &baseResult[S]{
		err: err,
	}
}

// NewWithResult creates a zero-value instance and applies settings.
// Parameters:
//   - settings: Configuration functions to apply
//
// Returns:
//   - *S: New configured instance
func NewWithResult[S any](settings []func(*S)) Result[S] {
	var zero S
	for _, setting := range settings {
		_ = apply(&zero, setting)
	}
	return NewValueResult(&zero)
}
