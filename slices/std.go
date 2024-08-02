package slices

import (
	"unicode"
	"unicode/utf8"
)

// A span is used to record a slice of s of the form s[start:end].
// The start index is inclusive and the end index is exclusive.
type span struct {
	start int
	end   int
}

// FieldsFunc splits the Array s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// Array is empty, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
func FieldsFunc[T ~[]S, S E](s T, f func(S) bool) []T {
	spans := make([]span, 0, 32)

	// Find the field start and end indices.
	// Doing this in a separate pass (rather than slicing the Array s
	// and collecting the result subArrays right away) is significantly
	// more efficient, possibly due to cache effects.
	start := -1 // valid span start if >= 0
	for end, rune := range s {
		if f(rune) {
			if start >= 0 {
				spans = append(spans, span{start, end})
				// Set start to a negative value.
				// Note: using -1 here consistently and reproducibly
				// slows down this code by a several percent on amd64.
				start = ^start
			}
		} else {
			if start < 0 {
				start = end
			}
		}
	}

	// Last field might end at EOF.
	if start >= 0 {
		spans = append(spans, span{start, len(s)})
	}

	// Create Arrays from recorded field indices.
	a := make([]T, len(spans))
	for i, span := range spans {
		a[i] = s[span.start:span.end]
	}

	return a
}

// Fields splits the Array s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of subArrays of s or an
// empty slice if s contains only white space.
func Fields[T ~[]S, S interface{ ~byte | ~rune }](s T) []T {
	// First count the fields.
	// This is an exact count if s is ASCII, otherwise it is an approximation.
	n := 0
	wasSpace := 1
	// setBits is used to track which bits are set in the bytes of s.
	setBits := S(0)
	for i := 0; i < len(s); i++ {
		r := s[i]
		setBits |= r
		isSpace := int(asciiSpace[r])
		n += wasSpace & ^isSpace
		wasSpace = isSpace
	}

	if setBits >= utf8.RuneSelf {
		// Some runes in the input Array are not ASCII.
		return FieldsFunc(s, IsSpace[S])
	}

	// ASCII fast path
	a := make([]T, n)
	na := 0
	fieldStart := 0
	i := 0
	// Skip spaces in the front of the input.
	for i < len(s) && asciiSpace[s[i]] != 0 {
		i++
	}
	fieldStart = i
	for i < len(s) {
		if asciiSpace[s[i]] == 0 {
			i++
			continue
		}
		a[na] = s[fieldStart:i]
		na++
		i++
		// Skip spaces in between fields.
		for i < len(s) && asciiSpace[s[i]] != 0 {
			i++
		}
		fieldStart = i
	}
	if fieldStart < len(s) { // Last field might end at EOF.
		a[na] = s[fieldStart:]
	}
	return a
}

func IsSpace[S interface{ ~byte | ~rune }](r S) bool {
	if r >= utf8.RuneSelf {
		return unicode.IsSpace(rune(r))
	}
	return asciiSpace[r] != 0
}
