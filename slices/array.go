package slices

import (
	"errors"
	"slices"

	"github.com/goexts/generic/types"
)

// E is comparable type of slice element
type E = comparable
type Slicer[T any] interface {
	~[]T
}

var (
	// ErrTooLarge is an error when number is too large than length
	ErrTooLarge = errors.New("slices.Array: number is too large than length")
	// ErrTooSmall is an error when number is too small than length
	ErrTooSmall = errors.New("slices.Array: number is too small than length")
	// ErrWrongIndex is an error when index is out of range
	ErrWrongIndex = errors.New("slices.Array: wrong index")
)

// Read returns a slice of the Array[S] s beginning at offset and length limit.
// If offset or limit is negative, it is treated as if it were zero.
func Read[T types.Slice[S], S E](arr T, offset int, limit int) T {
	if offset < 0 || limit < 0 {
		return nil
	}
	if len(arr) < offset+limit {
		return nil
	}
	return arr[offset : offset+limit]
}

// Append appends the element v to the end of Array[S] s.
func Append[T types.Slice[S], S E](arr T, v S) (T, int) {
	sz := len(arr)
	return append(arr, v), sz
}

// Equal reports whether a and b
// are the same length and contain the same runes.
// A nil argument is equivalent to an empty slice.
func Equal[T types.Slice[S], S E](a, b T) bool {
	return slices.Equal(a, b)
}

// HasPrefix tests whether the Array[S] s begins with prefix.
func HasPrefix[T types.Slice[S], S E](s, prefix T) bool {
	return len(s) >= len(prefix) && Equal[T](s[0:len(prefix)], prefix)
}

// HasSuffix tests whether the Array[S] s ends with suffix.
func HasSuffix[T types.Slice[S], S E](s, suffix T) bool {
	return len(s) >= len(suffix) && Equal(s[len(s)-len(suffix):], suffix)
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex[T types.Slice[S], S E](s, sep T) int {
	n := len(sep)
	switch {
	case n == 0:
		return len(s)
	case n == 1:
		return LastIndexArray(s, sep[0])
	case n == len(s):
		if Equal(s, sep) {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	last := len(s) - n
	if Equal(sep, s[last:]) {
		return last
	}
	for i := last - 1; i >= 0; i-- {
		if Equal(sep, s[i:i+n]) {
			return i
		}
	}
	return -1
}

// LastIndexArray returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexArray[T types.Slice[S], S E](s T, c S) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// Join concatenates the elements of its first argument to create a single Array[S]. The separator
// Array[S] sep is placed between elements in the resulting Array[S].
func Join[T types.Slice[S], S E](s []T, sep T) T {
	if len(s) == 0 {
		return T{}
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

// Repeat returns a new Array[S] consisting of count copies of the Array[S] s.
//
// It panics if count is negative or if
// the result of (len(s) * count) overflows.
func Repeat[T types.Slice[S], S E](b T, count int) T {
	if count == 0 {
		return T{}
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

// IndexArray returns the index of the first instance of the runes point
// r, or -1 if rune is not present in s.
func IndexArray[T types.Slice[S], S E](s T, r S) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index[T types.Slice[S], S E](s, sep T) int {
	n := len(sep)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return IndexArray(s, sep[0])
	case n == len(s):
		if Equal(sep, s) {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	c0 := sep[0]
	c1 := sep[1]
	i := 0
	t := len(s) - n + 1
	for i < t {
		if s[i] != c0 {
			o := IndexArray(s[i+1:t], c0)
			if o < 0 {
				return -1
			}
			i += o + 1
		}
		if s[i+1] == c1 && Equal(s[i:i+n], sep) {
			return i
		}
		i++
	}
	return -1
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
		i := Index(s, sub)
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

// Contains reports whether substr is within s.
func Contains[T types.Slice[S], S E](s, sub T) bool {
	return Index(s, sub) >= 0
}

// ContainsArray reports whether any Unicode code points in chars are within s.
func ContainsArray[T types.Slice[S], S E](s T, e S) bool {
	return IndexArray(s, e) >= 0
}

// explode splits s into a slice of UTF-8 Arrays,
// one Array per Unicode character up to a maximum of n (n < 0 means no limit).
// Invalid UTF-8 sequences become correct encodings of U+FFFD.
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
		m := Index(s, sep)
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

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func Cut[T types.Slice[S], S E](s, sep T) (before, after T, found bool) {
	if i := Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, nil, false
}

// InsertWith inserts v into s at the first index where fn(a, b) is true.
func InsertWith[T types.Slice[S], S E](s T, v S, fn func(a, b S) bool) T {
	pos := binarySearch(s, v, fn)

	// Create the result slice with the appropriate capacity.
	ret := make(T, 0, len(s)+1)
	if pos == -1 {
		return append(s, v)
	}
	// Append elements up to the insertion point.
	ret = append(ret, s[:pos]...)

	// Insert v.
	ret = append(ret, v)

	// Append the rest of the elements.
	ret = append(ret, s[pos:]...)

	return ret
}

// binarySearch performs a binary search to find the insertion point for v in s.
func binarySearch[S ~[]R, R E](s S, target R, cmp func(a, b R) bool) int {
	n := len(s)
	l, r := 0, n
	for l < r {
		mid := int(uint(l+r) >> 1)
		if !cmp(s[mid], target) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l < n && cmp(s[l], target) {
		return l
	}
	return -1
}

// RemoveWith removes the first index where fn(a, b) is true.
func RemoveWith[T types.Slice[S], S E](s T, fn func(a S) bool) T {
	ret := s[:0]
	for i := range s {
		if !fn(s[i]) {
			ret = append(ret, s[i])
		}
	}
	return ret
}

func CopyAt[T types.Slice[S], S E](s, t T, i int) T {
	if i < 0 {
		panic(ErrWrongIndex)
	}
	caps := cap(s)
	lent := len(s)
	if caps < lent+i {
		s = append(s, make([]S, lent+i-caps)...)
	}
	// copy the elements from s to t.
	copy(s[i:], t)
	return s
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
