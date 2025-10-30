/*
Package bytes provides a rich set of functions for the manipulation of byte slices.

This package is a generated adapter that mirrors the public API of the standard
Go library's `bytes` package. It offers a convenient way to access the rich
set of standard byte slice utilities within the generic context of this library.

For detailed information on the behavior of specific functions, please refer to
the official Go documentation for the `bytes` package.

# Usage

## Trimming and Cleaning

Remove unwanted characters from the beginning or end of a byte slice.

	data := []byte("  [INFO] message  ")

	// Trim leading and trailing whitespace.
	trimmed := bytes.TrimSpace(data)
	// trimmed is now []byte("[INFO] message")

	// Trim a specific prefix.
	noPrefix := bytes.TrimPrefix(trimmed, []byte("[INFO] "))
	// noPrefix is now []byte("message")

## Searching and Replacing

Find and replace subsequences within a byte slice.

	logLine := []byte("error: user not found, user_id=123")

	// Replace all occurrences of a subslice.
	updatedLog := bytes.ReplaceAll(logLine, []byte("error"), []byte("warning"))
	// updatedLog is []byte("warning: user not found, user_id=123")

	// Check for a prefix or suffix.
	hasPrefix := bytes.HasPrefix(logLine, []byte("error")) // true
	containsID := bytes.Contains(logLine, []byte("user_id")) // true

## Splitting and Joining

Divide a byte slice into parts or join parts together.

	csvRow := []byte("id,name,email")

	// Split the slice by a separator.
	fields := bytes.Split(csvRow, []byte(","))
	// fields is [][]byte{[]byte("id"), []byte("name"), []byte("email")}

	// Join a slice of byte slices with a separator.
	pathParts := [][]byte{[]byte("home"), []byte("user"), []byte("documents")}
	fullPath := bytes.Join(pathParts, []byte("/"))
	// fullPath is []byte("home/user/documents")
*/
package bytes
