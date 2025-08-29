package runes

import (
	"unicode"

	gslices "github.com/goexts/generic/slices"
)

//go:generate adptool .
//go:adapter:package golang.org/x/text/runes

// Runes is a type alias for []rune to provide methods.
type Runes []rune

// Read returns a slice of the Runes s beginning at offset and length limit.
func (r Runes) Read(offset int, limit int) []rune {
	if offset < 0 || limit < 0 {
		return nil
	}
	if len(r) < offset+limit {
		return nil
	}
	newSlice := make([]rune, limit)
	copy(newSlice, r[offset:offset+limit])
	return newSlice
}

// ReadString returns a string of the Runes s beginning at offset and length limit.
func (r Runes) ReadString(offset int, limit int) string {
	return string(r.Read(offset, limit))
}

// Index returns the index of the first instance of sub in r, or -1 if sub is not present in r.
func (r Runes) Index(sub []rune) int {
	return gslices.IndexSlice(r, sub)
}

// FindString returns the index of the first instance of s in r, or -1 if s is not present in r.
func (r Runes) FindString(s string) int {
	return r.Index([]rune(s))
}

// StringArray converts each rune to a string and returns a slice of strings.
func (r Runes) StringArray() []string {
	result := make([]string, 0, len(r))
	for i := range r {
		result = append(result, string(r[i]))
	}
	return result
}

// String converts the Runes slice to a string.
func (r Runes) String() string {
	return string(r)
}

// ToBytes converts the rune slice back to a UTF-8 encoded byte slice.
func (r Runes) ToBytes() []byte {
	return []byte(string(r))
}

// Trim returns a slice of the runes, with all leading and trailing runes contained in cutset removed.
func (r Runes) Trim(cutset string) []rune {
	cutsetRunes := []rune(cutset)
	start := 0
	for start < len(r) && gslices.Contains(cutsetRunes, r[start]) {
		start++
	}
	end := len(r)
	for end > start && gslices.Contains(cutsetRunes, r[end-1]) {
		end--
	}

	if start >= end {
		return nil
	}
	return r[start:end]
}

// TrimSpace returns a slice of the runes, with all leading and trailing white space removed.
func (r Runes) TrimSpace() []rune {
	start := 0
	for start < len(r) && unicode.IsSpace(r[start]) {
		start++
	}
	end := len(r)
	for end > start && unicode.IsSpace(r[end-1]) {
		end--
	}

	if start >= end {
		return nil
	}
	return r[start:end]
}

// TrimPrefix returns s without the provided leading prefix.
func (r Runes) TrimPrefix(prefix []rune) []rune {
	if r.HasPrefix(prefix) {
		return r[len(prefix):]
	}
	return r
}

// TrimSuffix returns s without the provided trailing suffix.
func (r Runes) TrimSuffix(suffix []rune) []rune {
	if r.HasSuffix(suffix) {
		return r[:len(r)-len(suffix)]
	}
	return r
}

// Replace returns a copy of the slice with the first n non-overlapping instances of old replaced by new.
func (r Runes) Replace(old, new []rune, n int) []rune {
	if len(old) == 0 || n == 0 {
		return gslices.Clone(r)
	}

	if n < 0 {
		n = gslices.Count(r, old)
	}

	if n == 0 {
		return gslices.Clone(r)
	}

	newLen := len(r) + n*(len(new)-len(old))
	if newLen < 0 {
		newLen = 0
	}
	result := make([]rune, 0, newLen)

	start := 0
	for i := 0; i < n; i++ {
		j := gslices.IndexSlice(r[start:], old)
		if j < 0 {
			break
		}
		result = append(result, r[start:start+j]...)
		result = append(result, new...)
		start += j + len(old)
	}
	result = append(result, r[start:]...)
	return result
}

// Contains reports whether sub is within r.
func (r Runes) Contains(sub []rune) bool {
	return gslices.IndexSlice(r, sub) != -1
}

// HasPrefix tests whether the Runes slice s begins with prefix.
func (r Runes) HasPrefix(prefix []rune) bool {
	return len(r) >= len(prefix) && gslices.Equal(r[0:len(prefix)], prefix)
}

// HasSuffix tests whether the Runes slice s ends with suffix.
func (r Runes) HasSuffix(suffix []rune) bool {
	return len(r) >= len(suffix) && gslices.Equal(r[len(r)-len(suffix):], suffix)
}

// FromString converts a string to a rune slice ([]rune).
// This is a convenience function that is equivalent to `[]rune(s)`.
func FromString(s string) Runes {
	return []rune(s)
}
