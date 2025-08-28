/*
Package strings provides a collection of functions for string manipulation.

This package is a generated adapter and mirrors the public API of the standard
Go library's `strings` package. It offers a convenient way to access the rich
set of standard string utilities.

For detailed information on the behavior of specific functions, please refer to
the official Go documentation for the `strings` package.

Example:

	addr := "[INFO] This is a log message."

	// Check for a prefix
	_ = strings.HasPrefix(addr, "[INFO]") // true

	// Trim the prefix
	msg := strings.TrimPrefix(addr, "[INFO] ")
	// msg is "This is a log message."
*/
package strings
