// 创建新文件并迁移代码
/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package result implements the unified result type for operations
package result

import (
	"github.com/goexts/generic/settings"
)

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

func With[S any](target *S, ss []func(*S)) Result[S] {
	_ = settings.Apply(target, ss)
	return NewValueResult(target)
}

// WithZero creates a zero-value instance and applies settings.
// Parameters:
//   - settings: Configuration functions to apply
//
// Returns:
//   - *S: New configured instance
func WithZero[S any](settings []func(*S)) Result[S] {
	var zero S
	return With(&zero, settings)
}

func WithZeroE[S any](settings []func(*S) error) Result[S] {
	var zero S
	return WithE(&zero, settings)
}

func WithE[S any](target *S, ss []func(*S) error) Result[S] {
	if _, err := settings.ApplyE(target, ss); err != nil {
		return NewErrorResult[S](err)
	}
	return NewValueResult[S](target)
}

func WithZeroMixed[S any](settings []any) Result[S] {
	var zero S
	return WithMixed(&zero, settings)
}

func WithMixed[S any](target *S, ss []any) Result[S] {
	_, err := settings.WithMixed(target, ss...)
	if err != nil {
		return NewErrorResult[S](err)
	}
	return NewValueResult[S](target)
}
