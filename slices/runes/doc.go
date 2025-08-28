/*
Package runes provides a rich set of functions for the manipulation of rune slices (`[]rune`).

This package is essential for correct, Unicode-aware text processing at the
code point level.

This package is a generated adapter and mirrors the public API of the Go
experimental package `golang.org/x/text/runes`. It offers a convenient way to
access these specialized utilities within the generic context of this library.

For detailed information on the behavior of specific functions and the underlying
Unicode algorithms, please refer to the official documentation for the
`golang.org/x/text/runes` package.

Example:

	text := []rune("  Hello, 世界!  ")

	// Trim whitespace using Unicode-aware functions
	trimmed := runes.TrimSpace(text)
	// trimmed is "Hello, 世界!"
*/
package runes
