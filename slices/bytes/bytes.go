package bytes

import (
	"slices"
)

//go:generate adptool bytes.go
//go:adapter:package bytes

// Bytes is a type alias for []byte to provide methods.
type Bytes []byte

// Read returns a slice of the Bytes s beginning at offset and length limit.
func (b Bytes) Read(offset int, limit int) []byte {
	if offset < 0 || limit < 0 {
		return nil
	}
	if len(b) < offset+limit {
		return nil
	}
	return slices.Clone(b[offset : offset+limit])
}

// ReadString returns a string of the Bytes s beginning at offset and length limit.
func (b Bytes) ReadString(offset int, limit int) string {
	return string(b.Read(offset, limit))
}

// Index returns the index of the first instance of sub in b, or -1 if sub is not present in b.
func (b Bytes) Index(sub []byte) int {
	return Index(b, sub)
}

// FindString returns the index of the first instance of s in b, or -1 if s is not present in b.
func (b Bytes) FindString(s string) int {
	return b.Index([]byte(s))
}

// String converts the Bytes slice to a string.
func (b Bytes) String() string {
	return string(b)
}

// Trim returns a slice of the bytes, with all leading and trailing bytes contained in cutset removed.
func (b Bytes) Trim(cutset string) []byte {
	return Trim(b, cutset)
}

// TrimSpace returns a slice of the bytes, with all leading and trailing white space removed.
func (b Bytes) TrimSpace() []byte {
	return TrimSpace(b)
}

// TrimPrefix returns b without the provided leading prefix.
func (b Bytes) TrimPrefix(prefix []byte) []byte {
	return TrimPrefix(b, prefix)
}

// TrimSuffix returns b without the provided trailing suffix.
func (b Bytes) TrimSuffix(suffix []byte) []byte {
	return TrimSuffix(b, suffix)
}

// Replace returns a copy of the slice with the first n non-overlapping instances of old replaced by replacement.
func (b Bytes) Replace(old, replacement []byte, n int) []byte {
	return Replace(b, old, replacement, n)
}

// Contains reports whether sub is within b.
func (b Bytes) Contains(sub []byte) bool {
	return Contains(b, sub)
}

// HasPrefix tests whether the byte slice b begins with prefix.
func (b Bytes) HasPrefix(prefix []byte) bool {
	return HasPrefix(b, prefix)
}

// HasSuffix tests whether the byte slice b ends with suffix.
func (b Bytes) HasSuffix(suffix []byte) bool {
	return HasSuffix(b, suffix)
}

// FromString converts a string to a Bytes slice.
func FromString(s string) Bytes {
	return []byte(s)
}
