package slices

import (
	"errors"
)

// E is comparable type of slice element
type E = comparable

// ErrTooLarge is an error when number is too large than length
var ErrTooLarge = errors.New("slices.Array: number is too large than length")

// Read returns a slice of the Array[S] s beginning at offset and length limit.
// If offset or limit is negative, it is treated as if it were zero.
func Read[T ~[]S, S E](arr T, offset int, limit int) T {
	if offset < 0 || limit < 0 {
		return nil
	}
	if len(arr) < offset+limit {
		return nil
	}
	return arr[offset : offset+limit]
}

// Equal reports whether a and b
// are the same length and contain the same runes.
// A nil argument is equivalent to an empty slice.
func Equal[T ~[]S, S E](a, b T) bool {
	lenA := len(a)
	lenB := len(b)
	if lenA == 0 || lenB == 0 {
		return lenA == 0 && lenB == 0
	}
	if lenA != lenB {
		return false
	}

	for i := 0; i < lenA; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// HasPrefix tests whether the Array[S] s begins with prefix.
func HasPrefix[T ~[]S, S E](s, prefix T) bool {
	return len(s) >= len(prefix) && Equal[T](s[0:len(prefix)], prefix)
}

// HasSuffix tests whether the Array[S] s ends with suffix.
func HasSuffix[T ~[]S, S E](s, suffix T) bool {
	return len(s) >= len(suffix) && Equal(s[len(s)-len(suffix):], suffix)
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex[T ~[]S, S E](s, sep T) int {
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
func LastIndexArray[T ~[]S, S E](s T, c S) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// Join concatenates the elements of its first argument to create a single Array[S]. The separator
// Array[S] sep is placed between elements in the resulting Array[S].
func Join[T ~[]S, S E](s []T, sep T) T {
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
func Repeat[T ~[]S, S E](b T, count int) T {
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

// indexFunc is the same as IndexFunc except that if
// truth==false, the sense of the predicate function is
// inverted.
func indexFunc[T ~[]S, S E](s T, f func(S) bool, truth bool) int {
	for i, r := range s {
		if f(r) == truth {
			return i
		}
	}
	return -1
}

// lastIndexFunc is the same as LastIndexFunc except that if
// truth==false, the sense of the predicate function is
// inverted.
func lastIndexFunc[T ~[]S, S E](s T, f func(S) bool, truth bool) int {
	for i := len(s); i > 0; {
		r := s[i-1]
		i--
		if f(r) == truth {
			return i
		}
	}
	return -1
}

// TrimPrefix returns s without the provided leading prefix Array.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix[T ~[]S, S E](s, prefix T) T {
	if HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

// TrimSuffix returns s without the provided trailing suffix Array.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix[T ~[]S, S E](s, suffix T) T {
	if HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

// IndexArray returns the index of the first instance of the runes point
// r, or -1 if rune is not present in s.
func IndexArray[T ~[]S, S E](s T, r S) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index[T ~[]S, S E](s, sep T) int {
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
func Count[T ~[]S, S E](s, sub T) int {
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
func CountArray[T ~[]S, S E](ss T, s S) int {
	n := 0
	for _, x := range ss {
		if x == s {
			n++
		}
	}
	return n
}

// Contains reports whether substr is within s.
func Contains[T ~[]S, S E](s, sub T) bool {
	return Index(s, sub) >= 0
}

// ContainsArray reports whether any Unicode code points in chars are within s.
func ContainsArray[T ~[]S, S E](s T, e S) bool {
	return IndexArray(s, e) >= 0
}

// explode splits s into a slice of UTF-8 Arrays,
// one Array per Unicode character up to a maximum of n (n < 0 means no limit).
// Invalid UTF-8 sequences become correct encodings of U+FFFD.
func explode[T ~[]S, S E](s T, n int) []T {
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
func genSplit[T ~[]S, S E](s, sep T, sepSave, n int) []T {
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

// SplitN slices s into sub arrays separated by sep and returns a slice of
// the sub arrays between those separators.
//
// The count determines the number of sub arrays to return:
//
//	n > 0: at most n sub arrays; the last subArray will be the unsplit remainder.
//	n == 0: the result is nil (zero sub arrays)
//	n < 0: all sub arrays
//
// Edge cases for s and sep (for example, empty Arrays) are handled
// as described in the documentation for Split.
//
// To split around the first instance of a separator, see Cut.
func SplitN[T ~[]S, S E](s, sep T, n int) []T { return genSplit(s, sep, 0, n) }

// SplitAfterN slices s into sub arrays after each instance of sep and
// returns a slice of those sub arrays.
//
// The count determines the number of sub arrays to return:
//
//	n > 0: at most n sub arrays; the last subArray will be the unsplit remainder.
//	n == 0: the result is nil (zero sub arrays)
//	n < 0: all sub arrays
//
// Edge cases for s and sep (for example, empty Arrays) are handled
// as described in the documentation for SplitAfter.
func SplitAfterN[T ~[]S, S E](s, sep T, n int) []T {
	return genSplit(s, sep, len(sep), n)
}

// Split slices s into all sub arrays separated by sep and returns a slice of
// the sub arrays between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
//
// To split around the first instance of a separator, see Cut.
func Split[T ~[]S, S E](s, sep T) []T { return genSplit(s, sep, 0, -1) }

// SplitAfter slices s into all sub arrays after each instance of sep and
// returns a slice of those sub arrays.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter[T ~[]S, S E](s, sep T) []T {
	return genSplit(s, sep, len(sep), -1)
}

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc[T ~[]S, S E](s T, f func(S) bool) int {
	return indexFunc(s, f, true)
}

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func LastIndexFunc[T ~[]S, S E](s T, f func(S) bool) int {
	return lastIndexFunc(s, f, true)
}

// Replace returns a copy of the Array s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the Array
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune Array.
// If n < 0, there is no limit on the number of replacements.
func Replace[T ~[]S, S E](s, old, new T, n int) T {
	m := 0
	if n != 0 {
		// Compute number of replacements.
		m = Count(s, old)
	}
	if m == 0 {
		// Just return a copy.
		return append(T(nil), s...)
	}
	if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	t := make(T, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				wid := len(s[start:])
				j += wid
			}
		} else {
			j += Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return t[0:w]
}

// ReplaceAll returns a copy of the Array s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the Array
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune Array.
func ReplaceAll[T ~[]S, S E](s, old, new T) T {
	return Replace(s, old, new, -1)
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func Cut[T ~[]S, S E](s, sep T) (before, after T, found bool) {
	if i := Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, nil, false
}

// TrimLeftFunc returns a slice of the Array s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc[T ~[]S, S E](s T, f func(S) bool) T {
	i := indexFunc(s, f, false)
	if i == -1 {
		return nil
	}
	return s[i:]
}

// TrimRightFunc returns a slice of the Array s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc[T ~[]S, S E](s T, f func(S) bool) T {
	i := lastIndexFunc(s, f, false)
	i++
	return s[0:i]
}

// TrimFunc returns a slice of the Array s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc[T ~[]S, S E](s T, f func(S) bool) T {
	return TrimRightFunc(TrimLeftFunc(s, f), f)
}

// Trim returns a sub slice of s by slicing off all leading and
// trailing UTF-8-encoded code points contained in cutset.
func Trim[T ~[]S, S E](s T, cutset T) T {
	if len(s) == 0 {
		// This is what we've historically done.
		return nil
	}

	switch len(cutset) {
	case 0:
		return s
	case 1:
		return trimLeftArray(trimRightArray(s, cutset[0]), cutset[0])
	default:
		return trimLeft(trimRight(s, cutset), cutset)
	}
}

// TrimLeft returns a sub slice of s by slicing off all leading
func TrimLeft[T ~[]S, S E](s T, cutset T) T {
	if len(s) == 0 {
		// This is what we've historically done.
		return nil
	}

	switch len(cutset) {
	case 0:
		return s
	case 1:
		return trimLeftArray(s, cutset[0])
	default:
		return trimLeft(s, cutset)

	}
}

func TrimRight[T ~[]S, S E](s T, cutset T) T {
	if len(s) == 0 {
		// This is what we've historically done.
		return nil
	}

	switch len(cutset) {
	case 0:
		return s
	case 1:
		return trimRightArray(s, cutset[0])
	default:
		return trimRight(s, cutset)
	}
}

func trimLeftArray[T ~[]S, S E](s T, c S) T {
	for len(s) > 0 && s[0] == c {
		s = s[1:]
	}
	return s
}

func trimLeft[T ~[]S, S E](s, cutset T) T {
	for len(s) > 0 {
		r, n := s[0], 1
		if !ContainsArray(cutset, r) {
			break
		}
		s = s[n:]
	}
	return s
}

func trimRightArray[T ~[]S, S E](s T, c S) T {
	for len(s) > 0 && s[len(s)-1] == c {
		s = s[:len(s)-1]
	}
	return s
}

func trimRight[T ~[]S, S E](s, cutset T) T {
	for len(s) > 0 {
		r, n := s[len(s)-1], 1
		if !ContainsArray(cutset, r) {
			break
		}
		s = s[:len(s)-n]
	}
	return s
}
