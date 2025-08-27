package slices

import (
	"bytes"
	"slices"
)

type Bytes []byte

func (r Bytes) Read(offset int, limit int) Bytes {
	if offset < 0 || limit < 0 {
		return nil
	}
	if offset+limit > len(r) {
		return nil
	}

	return slices.Clone(r[offset : offset+limit])
}

func (r Bytes) ReadString(offset int, limit int) string {
	return r.Read(offset, limit).String()
}

func (r Bytes) Index(sub []byte) int {
	return bytes.Index(r, sub)
}

func (r Bytes) FindString(s string) int {
	return r.Index([]byte(s))
}

// StringArray converts each byte in the slice to a separate string.
//
// WARNING: This method operates on individual bytes, not Unicode characters (runes).
// If the byte slice contains multi-byte UTF-8 characters, the result may be unexpected.
// For proper Unicode character handling, convert to Runes first using `ToRunes()`.
func (r Bytes) StringArray() []string {
	result := make([]string, 0, len(r))
	for i := range r {
		result = append(result, string(r[i]))
	}
	return result
}

func (r Bytes) String() string {
	return string(r)
}

// ToRunes converts the byte slice (assuming it is UTF-8 encoded) to a slice of runes.
func (r Bytes) ToRunes() Runes {
	return []rune(string(r))
}

func StringToBytes(s string) Bytes {
	return Bytes(s)
}
