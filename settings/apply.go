/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package settings implements the functions, types, and interfaces for the module.
package settings

// ApplyFunc is a ApplyFunc function for Apply
type ApplyFunc[S any] func(*S)

type ApplySetting[S any] interface {
	Apply(v *S)
}

type Setting[S any] interface {
	func(*S) | ApplyFunc[S]
}

func (s ApplyFunc[S]) Apply(v *S) {
	if s == nil {
		return
	}
	(s)(v)
}

func apply[S any](s *S, as ApplySetting[S]) {
	as.Apply(s)
}

// Apply is apply settings
func Apply[S any](s *S, fs []func(*S)) *S {
	if s == nil {
		return nil
	}
	for _, f := range fs {
		apply(s, ApplyFunc[S](f))
	}
	return s
}

// ApplyOr is an apply settings with defaults
func ApplyOr[S any](s *S, fs ...func(*S)) *S {
	return Apply(s, fs)
}

// ApplyOrZero is an apply settings with defaults
func ApplyOrZero[S any](fs ...func(*S)) *S {
	var val S
	return Apply(&val, fs)
}

// ApplyAny Applies a set of setting functions to a struct.
// These Settings functions can be ApplyFunc[S] or func(*S), or objects that implement the ApplySetting interface.
func ApplyAny[S any](s *S, fs []interface{}) *S {
	if s == nil {
		return nil
	}

	var as ApplySetting[S]
	for _, f := range fs {
		switch v := f.(type) {
		case ApplyFunc[S]:
			as = v
		case func(*S):
			as = ApplyFunc[S](v)
		case ApplySetting[S]:
			as = v
		default:
			panic("invalid apply setting type")
		}
		apply(s, as)
	}
	return s
}
