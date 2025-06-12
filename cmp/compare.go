// Package cmp implements the functions, types, and interfaces for the module.
package cmp

type Comparable interface {
	Equal(other any) bool
}

func Compare[T Comparable](a, b T) bool {
	return a.Equal(b)
}

func Equal[T comparable](a, b T) bool {
	return a == b
}
