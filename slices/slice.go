// Package slices implements the functions, types, and interfaces for the module.
package slices

import (
	"errors"
	"sort"

	"github.com/goexts/generic/types"
)

// E is comparable type of slice element
type E = comparable

var (
	// ErrWrongIndex is an error when index is out of range
	ErrWrongIndex = errors.New("slices.Array: wrong index")
)

// Append appends the element v to the end of Array[S] s.
func Append[T types.Slice[S], S any](arr T, v S) (T, int) {
	sz := len(arr)
	return append(arr, v), sz
}

// CopyAt copies the elements from t into s at the specified index.
// It panics if the index is negative. If the required length is greater
// than the length of s, s is grown to accommodate the new elements.
func CopyAt[T types.Slice[S], S any](s, t T, i int) T {
	if i < 0 {
		panic(ErrWrongIndex) // Or return an error, panic is consistent with stdlib
	}

	requiredLen := i + len(t)
	if len(s) < requiredLen {
		// Grow the slice to the required length, filling with zero values.
		s = append(s, make(T, requiredLen-len(s))...)
	}

	// Copy the elements from t into s at the specified index.
	copy(s[i:], t)
	return s
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty Array, Count returns 1 + the number of Unicode code points in s.
func Count[T types.Slice[S], S E](s, sub T) int {
	// special case
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

// CountArray counts the number of non-overlapping instances of c in s.
func CountArray[T types.Slice[S], S E](ss T, s S) int {
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
// If sep does not appear in s, cut returns s, "", false.
func Cut[T types.Slice[S], S E](s, sep T) (before, after T, found bool) {
	if i := IndexSlice(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, nil, false
}

// Filter returns a new slice containing all elements of s for which f(s) is true.
func Filter[T types.Slice[S], S any](s T, f func(S) bool) T {
	if s == nil {
		return make(T, 0)
	}
	result := make(T, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// IndexSlice returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func IndexSlice[T types.Slice[S], S E](s, substr T) int {
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
func InsertWith[T types.Slice[S], S any](s T, v S, fn func(a, b S) bool) T {
	// Assumes s is sorted according to fn.
	// sort.Search finds the first index `i` where the function is true.
	pos := sort.Search(len(s), func(i int) bool { return fn(s[i], v) })

	// Efficiently insert the element.
	s = append(s, *new(S))   // Grow slice by one (zero value).
	copy(s[pos+1:], s[pos:]) // Shift elements from pos onwards to the right.
	s[pos] = v               // Insert the new element at pos.
	return s
}

// Join concatenates the elements of its first argument to create a single Array[S]. The separator
// Array[S] sep is placed between elements in the resulting Array[S].
func Join[T types.Slice[S], S any](s []T, sep T) T {
	if len(s) == 0 {
		return make(T, 0)
	}
	if len(s) == 1 {
		// Just return a copy.
		return append(T(nil), s[0]...)
	}
	n := len(sep) * (len(s) - 1)
	for _, v := range s {
		n += len(v)
	}

	b := make(T, n)
	bp := copy(b, s[0])
	for _, v := range s[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], v)
	}
	return b
}

// LastIndexSlice returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndexSlice[T types.Slice[S], S E](s, sep T) int {
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

// Map transforms a slice of one type to a slice of another type by applying a function to each element.
func Map[S, T any](s []S, f func(S) T) []T {
	if s == nil {
		return make([]T, 0)
	}
	result := make([]T, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

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

// Read returns a slice of the Array[S] s beginning at offset and length limit.
// If offset or limit is negative, it is treated as if it were zero.
func Read[T types.Slice[S], S any](arr T, offset int, limit int) T {
	if offset < 0 || limit < 0 {
		return nil
	}
	if len(arr) < offset+limit {
		return nil
	}
	return clone(arr[offset : offset+limit])
}

// Reduce aggregates all elements of a slice into a single value by applying a function.
// It iterates through the slice, applying the function 'f' to an accumulator and the current element.
func Reduce[S, T any](s []S, initial T, f func(T, S) T) T {
	accumulator := initial
	for _, v := range s {
		accumulator = f(accumulator, v)
	}
	return accumulator
}

// RemoveWith removes the first index where fn(a, b) is true.
func RemoveWith[T types.Slice[S], S any](s T, fn func(a S) bool) T {
	ret := s[:0]
	for i := range s {
		if !fn(s[i]) {
			ret = append(ret, s[i])
		}
	}
	return ret
}

// Repeat returns a new Array[S] consisting of count copies of the Array[S] s.
//
// It panics if count is negative or if
// the result of (len(s) * count) overflows.
func Repeat[T types.Slice[S], S any](b T, count int) T {
	if count == 0 {
		return make(T, 0)
	}
	// Since we cannot return an error on overflow,
	// we should panic if the repeat will generate
	// an overflow.
	// See Issue golang.org/issue/16237.
	if count < 0 {
		panic("bytes: negative Repeat count")
	} else if len(b)*count/count != len(b) {
		panic("bytes: Repeat count causes overflow")
	}

	nb := make(T, len(b)*count)
	bp := copy(nb, b)
	for bp < len(nb) {
		copy(nb[bp:], nb[:bp])
		bp *= 2
	}
	return nb
}

// Split slices s into all subslices separated by sep and returns a slice of
// the subslices between those separators.
func Split[T types.Slice[S], S E](s, sep T) []T {
	if s == nil {
		return nil
	}
	return genSplit(s, sep, 0, -1)
}

func Transform[TS types.Slice[S], S any, T any](s TS, f func(S) (T, bool)) []T {
	if s == nil {
		return make([]T, 0)
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
// The order of the first occurrence of each element is preserved.
func Unique[T types.Slice[S], S E](s T) T {
	if len(s) == 0 {
		return make(T, 0)
	}
	seen := make(map[S]struct{})
	result := make(T, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func clone[T types.Slice[S], S any](s T) T {
	if s == nil {
		return nil
	}
	return append(T(nil), s...)
}

// equal reports whether two slices are equal: the same length and all
// elements equal.
func equal[T types.Slice[S], S E](s1, s2 T) bool {
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

// explode splits s into a slice of slices, each of length 1,
// up to a maximum of n (n < 0 means no limit).
func explode[T types.Slice[S], S E](s T, n int) []T {
	l := len(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([]T, n)
	for i := 0; i < n-1; i++ {
		a[i] = s[:1]
		s = s[1:]
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}

// Generic split: splits after each instance of sep,
// including sepSave bytes of sep in the sub arrays.
func genSplit[T types.Slice[S], S E](s, sep T, sepSave, n int) []T {
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

	a := make([]T, n)
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
func indexElement[T types.Slice[S], S E](s T, e S) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

// lastIndexElement returns the index of the last instance of e in s, or -1 if not found.
func lastIndexElement[T types.Slice[S], S E](s T, e S) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == e {
			return i
		}
	}
	return -1
}
