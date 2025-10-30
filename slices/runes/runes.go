package runes

import (
	"slices"
	"unicode"
)

//go:generate adptool runes.go
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
	return Index(r, sub)
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
	for start < len(r) && slices.Contains(cutsetRunes, r[start]) {
		start++
	}
	end := len(r)
	for end > start && slices.Contains(cutsetRunes, r[end-1]) {
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

// Replace returns a copy of the slice with the first n non-overlapping instances of old replaced by replacement.
func (r Runes) Replace(old, replacement []rune, n int) []rune {
	if len(old) == 0 || n == 0 {
		return slices.Clone(r)
	}

	if n < 0 {
		n = Count(r, old)
	}

	if n == 0 {
		return slices.Clone(r)
	}

	newLen := len(r) + n*(len(replacement)-len(old))
	if newLen < 0 {
		newLen = 0
	}
	result := make([]rune, 0, newLen)

	start := 0
	for i := 0; i < n; i++ {
		j := Index(r[start:], old)
		if j < 0 {
			break
		}
		result = append(result, r[start:start+j]...)
		result = append(result, replacement...)
		start += j + len(old)
	}
	result = append(result, r[start:]...)
	return result
}

// Contains reports whether sub is within r.
func (r Runes) Contains(sub []rune) bool {
	return Index(r, sub) != -1
}

// HasPrefix tests whether the Runes slice s begins with prefix.
func (r Runes) HasPrefix(prefix []rune) bool {
	return len(r) >= len(prefix) && slices.Equal(r[0:len(prefix)], prefix)
}

// HasSuffix tests whether the Runes slice s ends with suffix.
func (r Runes) HasSuffix(suffix []rune) bool {
	return len(r) >= len(suffix) && slices.Equal(r[len(r)-len(suffix):], suffix)
}

// Clone returns a copy of the Runes slice.
func (r Runes) Clone() Runes {
	return slices.Clone(r)
}

// FromString converts a string to a rune slice ([]rune).
// This is a convenience function that is equivalent to `[]rune(s)`.
func FromString(s string) Runes {
	return []rune(s)
}

// Index returns the index of the first instance of sub in r, or -1 if sub is not present in r.
func Index(r, sub []rune) int {
	// Follow bytes.Index semantics: empty pattern returns 0
	if len(sub) == 0 {
		return 0
	}
	// Optimize single-rune search
	if len(sub) == 1 {
		needle := sub[0]
		for i := 0; i < len(r); i++ {
			if r[i] == needle {
				return i
			}
		}
		return -1
	}

	// KMP (Knuth-Morris-Pratt) for general case
	// Build longest prefix-suffix (lps) array for pattern `sub`
	lps := make([]int, len(sub))
	for i, j := 1, 0; i < len(sub); {
		switch {
		case sub[i] == sub[j]:
			j++
			lps[i] = j
			i++
		case j != 0:
			j = lps[j-1]
		default:
			lps[i] = 0
			i++
		}
	}

	// Search using KMP
	for i, j := 0, 0; i < len(r); {
		switch {
		case r[i] == sub[j]:
			i++
			j++
			if j == len(sub) {
				return i - j
			}
		case j != 0:
			j = lps[j-1]
		default:
			i++
		}
	}
	return -1
}

// Count counts the number of non-overlapping instances of sub in r.
func Count(r, sub []rune) int {
	// Follow bytes.Count semantics: non-overlapping occurrences; empty pattern yields len(r)+1
	if len(sub) == 0 {
		return len(r) + 1
	}
	// Optimize single-rune counting
	if len(sub) == 1 {
		c := 0
		needle := sub[0]
		for i := 0; i < len(r); i++ {
			if r[i] == needle {
				c++
			}
		}
		return c
	}

	count := 0
	start := 0
	for {
		j := Index(r[start:], sub)
		if j < 0 {
			break
		}
		count++
		start += j + len(sub)
	}
	return count
}
