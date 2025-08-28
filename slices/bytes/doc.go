/*
Package bytes provides a rich set of functions for the manipulation of byte slices.

This package is a generated adapter and mirrors the public API of the standard
Go library's `bytes` package. It offers a convenient way to access the rich
set of standard byte slice utilities within the generic context of this library.

For detailed information on the behavior of specific functions, please refer to
the official Go documentation for the `bytes` package.

Example:

	data := []byte("  [INFO] message  ")

	// Trim whitespace
	trimmed := bytes.TrimSpace(data)
	// trimmed is "[INFO] message"

	// Check for a prefix
	_ = bytes.HasPrefix(trimmed, []byte("[INFO]")) // true
*/
package bytes
