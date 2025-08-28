package runes

//go:generate adptool .
//go:adapter:package golang.org/x/text/runes

// FromString converts a string to a rune slice ([]rune).
// This is a convenience function that is equivalent to `[]rune(s)`.
func FromString(s string) []rune {
	return []rune(s)
}
