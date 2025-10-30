// Package slices implements the functions, types, and interfaces for the module.
package slices

import (
	"errors"
	"sort"
)

// E is a shorthand for the comparable constraint.
type E = comparable

var (
	// ErrWrongIndex is an error returned when an index is out of range.
	ErrWrongIndex = errors.New("slices: index out of range")
)

// Append appends the element v to the end of the slice s.
func Append[S any](arr []S, v S) ([]S, int) {
	sz := len(arr)
	return append(arr, v), sz
}

// CopyAt copies the elements from t into s at the specified index.
// It panics if the index is negative. If the required length is greater
// than the length of s, s is grown to accommodate the new elements.
func CopyAt[S any](s, t []S, i int) []S {
	if i < 0 {
		panic(ErrWrongIndex)
	}

	requiredLen := i + len(t)
	if len(s) < requiredLen {
		s = append(s, make([]S, requiredLen-len(s))...)
	}

	copy(s[i:], t)
	return s
}

// Count counts the number of non-overlapping instances of substr in s.
func Count[S E](s, sub []S) int {
	if len(sub) == 0 {
		return len(s) + 1
	}
	if len(sub) == 1 {
		return CountArray(s, sub[0])
	}
	n := 0
	for {
		i := IndexSlice(s, sub)
		if i == -1 {
			return n
		}
		n++
		s = s[i+len(sub):]
	}
}

// CountArray counts the number of occurrences of c in s.
func CountArray[S E](ss []S, s S) int {
	n := 0
	for _, x := range ss {
		if x == s {
			n++
		}
	}
	return n
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, nil, false.
func Cut[S E](s, sep []S) (before, after []S, found bool) {
	if i := IndexSlice(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, nil, false
}

// Filter returns a new slice containing all elements of s for which f(s) is true.
func Filter[S any](s []S, f func(S) bool) []S {
	if s == nil {
		return nil
	}
	result := make([]S, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterIncluded returns a new slice containing all elements of s that are present in includes.
func FilterIncluded[S comparable](s []S, includes []S) []S {
	if len(includes) == 0 {
		return make([]S, 0)
	}
	if len(includes) <= 20 {
		return Filter(s, func(e S) bool {
			for _, v := range includes {
				if v == e {
					return true
				}
			}
			return false
		})
	}
	includeSet := make(map[S]struct{}, len(includes))
	for _, v := range includes {
		includeSet[v] = struct{}{}
	}
	return Filter(s, func(e S) bool {
		_, exists := includeSet[e]
		return exists
	})
}

// FilterExcluded returns a new slice containing all elements of s that are not present in excludes.
func FilterExcluded[S comparable](s []S, excludes []S) []S {
	if len(excludes) == 0 {
		return s
	}
	if len(excludes) <= 20 {
		return Filter(s, func(e S) bool {
			for _, v := range excludes {
				if v == e {
					return false
				}
			}
			return true
		})
	}
	excludeSet := make(map[S]struct{}, len(excludes))
	for _, v := range excludes {
		excludeSet[v] = struct{}{}
	}
	return Filter(s, func(e S) bool {
		_, exists := excludeSet[e]
		return !exists
	})
}

// IndexSlice returns the index of the first instance of substr in s, or -1 if not present.
func IndexSlice[S E](s, substr []S) int {
	n := len(substr)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return indexElement(s, substr[0])
	case n == len(s):
		if equal(s, substr) {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	for i := 0; i <= len(s)-n; i++ {
		if equal(s[i:i+n], substr) {
			return i
		}
	}
	return -1
}

// InsertWith inserts v into s at the first index where fn(a, b) is true.
func InsertWith[S any](s []S, v S, fn func(a, b S) bool) []S {
	pos := sort.Search(len(s), func(i int) bool { return fn(s[i], v) })
	s = append(s, *new(S))
	copy(s[pos+1:], s[pos:])
	s[pos] = v
	return s
}

// Join concatenates the elements of s to create a single slice.
// The separator sep is placed between elements in the resulting slice.
func Join[S any](s [][]S, sep []S) []S {
	if len(s) == 0 {
		return make([]S, 0)
	}
	if len(s) == 1 {
		return clone(s[0])
	}
	n := len(sep) * (len(s) - 1)
	for _, v := range s {
		n += len(v)
	}

	b := make([]S, n)
	bp := copy(b, s[0])
	for _, v := range s[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], v)
	}
	return b
}

// LastIndexSlice returns the index of the last instance of sep in s, or -1 if not present.
func LastIndexSlice[S E](s, sep []S) int {
	n := len(sep)
	switch {
	case n == 0:
		return len(s)
	case n == 1:
		return lastIndexElement(s, sep[0])
	case n == len(s):
		if equal(s, sep) {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	for i := len(s) - n; i >= 0; i-- {
		if equal(s[i:i+n], sep) {
			return i
		}
	}
	return -1
}

// Map transforms a slice of one type to a slice of another type using a mapping function.
func Map[S, T any](s []S, f func(S) T) []T {
	if s == nil {
		return nil
	}
	result := make([]T, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

// OverWithError returns an iterator function for a slice that may have an associated error.
func OverWithError[S any](s []S, err error) func(func(int, S) bool) {
	return func(yield func(int, S) bool) {
		if err != nil || len(s) == 0 {
			return
		}
		for i := range s {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

// Read returns a slice of s beginning at offset and with a length of limit.
func Read[S any](arr []S, offset int, limit int) []S {
	if offset < 0 || limit < 0 {
		return nil
	}
	if len(arr) < offset+limit {
		return nil
	}
	return clone(arr[offset : offset+limit])
}

// Reduce aggregates all elements of a slice into a single value by applying a function.
func Reduce[S, T any](s []S, initial T, f func(T, S) T) T {
	accumulator := initial
	for _, v := range s {
		accumulator = f(accumulator, v)
	}
	return accumulator
}

// RemoveWith removes elements from a slice based on a predicate function.
func RemoveWith[S any](s []S, fn func(a S) bool) []S {
	ret := s[:0]
	for i := range s {
		if !fn(s[i]) {
			ret = append(ret, s[i])
		}
	}
	return ret
}

// Repeat returns a new slice consisting of count copies of the slice s.
func Repeat[S any](b []S, count int) []S {
	if count == 0 {
		return make([]S, 0)
	}
	if count < 0 {
		panic("slices: negative Repeat count")
	} else if len(b)*count/count != len(b) {
		panic("slices: Repeat count causes overflow")
	}

	nb := make([]S, len(b)*count)
	bp := copy(nb, b)
	for bp < len(nb) {
		copy(nb[bp:], nb[:bp])
		bp *= 2
	}
	return nb
}

// Split slices s into all subslices separated by sep.
func Split[S E](s, sep []S) [][]S {
	if s == nil {
		return nil
	}
	return genSplit(s, sep, 0, -1)
}

// Transform combines mapping and filtering a slice.
func Transform[S any, T any](s []S, f func(S) (T, bool)) []T {
	if s == nil {
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

// Unique returns a new slice with duplicate elements removed.
func Unique[S E](s []S) []S {
	if len(s) == 0 {
		return make([]S, 0)
	}
	seen := make(map[S]struct{})
	result := make([]S, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// clone returns a shallow copy of the slice.
func clone[S any](s []S) []S {
	if s == nil {
		return nil
	}
	return append([]S(nil), s...)
}

// equal reports whether two slices are equal.
func equal[S E](s1, s2 []S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// explode splits s into a slice of slices of length 1.
func explode[S E](s []S, n int) [][]S {
	l := len(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([][]S, n)
	for i := 0; i < n-1; i++ {
		a[i] = s[:1]
		s = s[1:]
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}

// genSplit is the generic split implementation.
func genSplit[S E](s, sep []S, sepSave, n int) [][]S {
	if n == 0 {
		return nil
	}
	if len(sep) == 0 {
		return explode(s, n)
	}
	if n < 0 {
		n = Count(s, sep) + 1
	}
	if n > len(s)+1 {
		n = len(s) + 1
	}

	a := make([][]S, n)
	n--
	i := 0
	for i < n {
		m := IndexSlice(s, sep)
		if m < 0 {
			break
		}
		a[i] = s[: m+sepSave : m+sepSave]
		s = s[m+len(sep):]
		i++
	}
	a[i] = s
	return a[:i+1]
}

// indexElement returns the index of the first instance of e in s, or -1 if not found.
func indexElement[S E](s []S, e S) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

// lastIndexElement returns the index of the last instance of e in s, or -1 if not found.
func lastIndexElement[S E](s []S, e S) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == e {
			return i
		}
	}
	return -1
}
