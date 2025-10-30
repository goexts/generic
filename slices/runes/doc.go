/*
Package runes provides a rich set of functions for the manipulation of rune slices (`[]rune`).

This package is essential for correct, Unicode-aware text processing at the
code point level. It is a generated adapter that mirrors the public API of the
Go experimental package `golang.org/x/text/runes`, offering convenient access
to its specialized utilities.

For detailed information on the behavior of specific functions and the underlying
Unicode algorithms, please refer to the official documentation for the
`golang.org/x/text/runes` package.

# Usage

Working with `[]rune` is crucial when characters may be multi-byte, as simple
indexing on a `string` can split a Unicode code point.

## Unicode-Aware Trimming

Standard trimming functions might not handle all Unicode space characters correctly.

	text := []rune("\u2003Hello, 世界!\u2003") // \u2003 is a wide space

	// TrimSpace correctly removes various Unicode space characters.
	trimmed := runes.TrimSpace(text)
	// trimmed is []rune("Hello, 世界!")

## Filtering by Unicode Properties

You can filter runes based on their Unicode properties, such as being a letter,
number, or symbol.

	mixed := []rune("Go-1.18 is awesome! 🎉")

	// Filter to get only the letters.
	letters := runes.Filter(mixed, unicode.IsLetter)
	// letters is []rune("Goisawesome")

	// Filter to get only the punctuation and symbols.
	punctuation := runes.Filter(mixed, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	})
	// punctuation is []rune("-!🎉")

## Transforming Runes

`Map` allows you to transform each rune in a slice. This is useful for tasks
like case conversion.

	message := []rune("Hello, World")

	// Convert the entire slice to uppercase.
	upper := runes.Map(unicode.ToUpper, message)
	// upper is []rune("HELLO, WORLD")
*/
package runes
