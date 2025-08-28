package bytes

//go:generate adptool .
//go:adapter:package bytes

// FromString converts a string to a byte slice ([]byte).
// This is a convenience function that is equivalent to `[]byte(s)`.
func FromString(s string) []byte {
	return []byte(s)
}
