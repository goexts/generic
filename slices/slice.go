// Package slices implements the functions, types, and interfaces for the module.
package slices

import (
	"github.com/goexts/generic/types"
)

func Transform[TS types.Slice[S], S any, T any](s TS, f func(S) (T, bool)) []T {
	if len(s) == 0 {
		return nil
	}
	tt := make([]T, 0, len(s))
	for _, sv := range s {
		if t, ok := f(sv); ok {
			tt = append(tt, t)
		}
	}
	return tt
}
