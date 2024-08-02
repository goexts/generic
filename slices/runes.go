package slices

import (
	"slices"
	"unicode"
	"unicode/utf8"
)

type Runes []rune

func (r Runes) Read(offset int, limit int) Runes {
	if offset < 0 || limit < 0 {
		return nil
	}
	if offset+limit > len(r) {
		return nil
	}

	return slices.Clone(r[offset : offset+limit])
}

func (r Runes) ReadString(offset int, limit int) string {
	return r.Read(offset, limit).String()
}

func (r Runes) Index(sub []rune) int {
	return Index(r, sub)
}

func (r Runes) FindString(s string) int {
	return r.Index([]rune(s))
}

func (r Runes) StringArray() []string {
	var result []string
	for i := range r {
		result = append(result, string(r[i]))
	}
	return result
}

func (r Runes) String() string {
	return string(r)
}

func (r Runes) Map(mapping func(rune) rune) Runes {
	return MapRune(mapping, r)
}

func StringToRunes(s string) Runes {
	return Runes(s)
}

// MapRune returns a copy of the []rune s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the []rune with no replacement.
func MapRune(mapping func(rune) rune, s []rune) []rune {
	// In the worst case, the slice can grow when mapped, making
	// things unpleasant. But it's so rare we barge in assuming it's
	// fine. It could also shrink but that falls out naturally.
	maxrunes := len(s) // length of b
	nrunes := 0        // number of bytes encoded in b
	b := make([]rune, maxrunes)
	for i := 0; i < len(s); {
		wid := 1
		r := s[i]
		r = mapping(r)
		if r >= 0 {
			rl := utf8.RuneLen(r)
			if rl < 0 {
				rl = len(string(utf8.RuneError))
			}
			if nrunes+rl > maxrunes {
				// Grow the buffer.
				maxrunes = maxrunes*2 + utf8.UTFMax
				nb := make([]rune, maxrunes)
				copy(nb, b[0:nrunes])
				b = nb
			}
			nrunes++
		}
		i += wid
	}
	return b[0:nrunes]
}

// ToUpperRune returns s with all Unicode letters mapped to their upper case.
func ToUpperRune(s []rune) []rune {
	hasLower := true
	for i := 0; i < len(s); i++ {
		c := s[i]
		hasLower = 'a' <= c && c <= 'z'
	}

	if !hasLower {
		// Just return a copy.
		return append([]rune(""), s...)
	}
	b := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
		}
		b[i] = c
	}
	return b
}

// TrimSpaceRune returns a slice of the []rune s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpaceRune(s []rune) []rune {
	// Fast path for ASCII: look for the first ASCII non-space byte
	start := 0
	for ; start < len(s); start++ {
		c := s[start]
		if c >= utf8.RuneSelf {
			// If we run into a non-ASCII byte, fall back to the
			// slower unicode-aware method on the remaining bytes
			return TrimFunc(s[start:], unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}

	// Now look for the first ASCII non-space byte from the end
	stop := len(s)
	for ; stop > start; stop-- {
		c := s[stop-1]
		if c >= utf8.RuneSelf {
			// start has been already trimmed above, should trim end only
			return TrimRightFunc(s[start:stop], unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}

	// At this point s[start:stop] starts and ends with an ASCII
	// non-space bytes, so we're done. Non-ASCII cases have already
	// been handled above.
	return s[start:stop]
}

// EqualFoldRune reports whether s and t, interpreted as UTF-8 []runes,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
func EqualFoldRune(s, t []rune) bool {
	for len(s) != 0 && len(t) != 0 {
		// Extract first rune from each []rune.
		var sr, tr rune
		sr, s = s[0], s[1:]
		tr, t = t[0], t[1:]
		// If they match, keep going; if not, return false.

		// Easy case.
		if tr == sr {
			continue
		}

		// Make sr < tr to simplify what follows.
		if tr < sr {
			tr, sr = sr, tr
		}
		// Fast check for ASCII.
		if tr < utf8.RuneSelf {
			// ASCII only, sr/tr must be upper/lower case
			if 'A' <= sr && sr <= 'Z' && tr == sr+'a'-'A' {
				continue
			}
			return false
		}

		// General case. SimpleFold(x) returns the next equivalent rune > x
		// or wraps around to smaller values.
		r := unicode.SimpleFold(sr)
		for r != sr && r < tr {
			r = unicode.SimpleFold(r)
		}
		if r == tr {
			continue
		}
		return false
	}

	// One []rune is empty. Are both?
	return len(s) == len(t)
}

// IndexByteRune returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByteRune(s []rune, c byte) int {
	return IndexRune(s, rune(c))
}

// IndexRune returns the index of the first instance of the runes point
// r, or -1 if rune is not present in s.
func IndexRune(s []rune, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

// containsRune is a simplified version of strings.ContainsRune
// to avoid importing the strings package.
// We avoid bytes.ContainsRune to avoid allocating a temporary copy of s.
func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s []rune, r rune) bool {
	return IndexRune(s, r) >= 0
}

// ToLowerRune returns s with all Unicode letters mapped to their lower case.
func ToLowerRune(s []rune) []rune {
	hasUpper := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		hasUpper = hasUpper || ('A' <= c && c <= 'Z')
	}

	if !hasUpper {
		return append([]rune(""), s...)
	}
	b := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return b
}

// ToTitleRune returns a copy of the []rune s with all Unicode letters mapped to
// their Unicode title case.
func ToTitleRune(s []rune) []rune { return MapRune(unicode.ToTitle, s) }

// ToUpperSpecialRune returns a copy of the []rune s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
func ToUpperSpecialRune(c unicode.SpecialCase, s []rune) []rune {
	return MapRune(c.ToUpper, s)
}

// ToLowerSpecialRune returns a copy of the []rune s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
func ToLowerSpecialRune(c unicode.SpecialCase, s []rune) []rune {
	return MapRune(c.ToLower, s)
}

// ToTitleSpecialRune returns a copy of the []rune s with all Unicode letters mapped to their
// Unicode title case, giving priority to the special casing rules.
func ToTitleSpecialRune(c unicode.SpecialCase, s []rune) []rune {
	return MapRune(c.ToTitle, s)
}

// ToValidUTF8Rune returns a copy of the []rune s with each run of invalid UTF-8 byte sequences
// replaced by the replacement []rune, which may be empty.
func ToValidUTF8Rune(s, replacement []rune) []rune {
	b := make([]rune, 0, len(s)+len(replacement))
	invalid := false // previous byte was from an invalid UTF-8 sequence
	for i := 0; i < len(s); {
		c := s[i]
		if c < utf8.RuneSelf {
			i++
			invalid = false
			b = append(b, c)
			continue
		}
		i++
		if !invalid {
			invalid = true
			b = append(b, replacement...)
		}
		continue
	}
	return b
}

// isSeparatorRune reports whether the rune could mark a word boundary.
// TODO: update when package unicode captures more of the properties.
func isSeparatorRune(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}

// TitleRune returns a copy of the []rune s with all Unicode letters that begin words
// mapped to their Unicode title case.
//
// Deprecated: The rule TitleRune uses for word boundaries does not handle Unicode
// punctuation properly. Use golang.org/x/text/cases instead.
func TitleRune(s []rune) []rune {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on MapRune scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return MapRune(
		func(r rune) rune {
			if isSeparatorRune(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// asciiSet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
// starting with the least-significant bit of the lowest word to the
// most-significant bit of the highest word, map to the full range of all
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
// ensuring that any non-ASCII character will be reported as not in the set.
// This allocates a total of 32 bytes even though the upper half
// is unused to avoid bounds checks in asciiSet.contains.
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
// characters in chars are ASCII.
func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c/32] |= 1 << (c % 32)
	}
	return as, true
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c rune) bool {
	return (as[c/32] & (1 << (c % 32))) != 0
}

// TrimRightRune returns a slice of the []rune s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use TrimSuffix instead.
func TrimRightRune(s []rune, cutset string) []rune {
	if len(s) == 0 || cutset == "" {
		return s
	}

	if len(cutset) == 1 && cutset[0] < utf8.RuneSelf {
		return trimRightArray(s, []rune(cutset)[0])
	}

	if as, ok := makeASCIISet(cutset); ok {
		return trimRightASCII(s, &as)
	}

	return trimRightUnicode(s, cutset)
}

func trimRightASCII(s []rune, as *asciiSet) []rune {
	for len(s) > 0 {
		if !as.contains(s[len(s)-1]) {
			break
		}
		s = s[:len(s)-1]
	}
	return s
}

func trimRightUnicode(s []rune, cutset string) []rune {
	for len(s) > 0 {
		r, n := s[len(s)-1], 1
		if !containsRune(cutset, r) {
			break
		}
		s = s[:len(s)-n]
	}
	return s
}
