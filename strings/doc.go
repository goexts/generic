/*
Package strings provides a rich collection of functions for string manipulation.

This package is a generated adapter that mirrors the public API of the standard
Go library's `strings` package. It offers a convenient way to access the rich
set of standard string utilities.

For detailed information on the behavior of specific functions, please refer to
the official Go documentation for the `strings` package.

# Usage

## Searching and Replacing

Find, check for, and replace substrings.

	logLine := "[ERROR] user authentication failed: invalid token"

	// Check for substrings.
	isError := strings.HasPrefix(logLine, "[ERROR]") // true
	hasToken := strings.Contains(logLine, "token")      // true

	// Replace part of the string.
	warningLine := strings.Replace(logLine, "[ERROR]", "[WARNING]", 1)
	// warningLine is "[WARNING] user authentication failed: invalid token"

## Splitting and Joining

Split strings into a slice or join a slice of strings into one.

	// Splitting a string by a separator.
	path := "/usr/local/bin"
	parts := strings.Split(path, "/")
	// parts is []string{"", "usr", "local", "bin"}

	// Joining a slice of strings with a separator.
	csvFields := []string{"user", "email", "id"}
	csvHeader := strings.Join(csvFields, ",")
	// csvHeader is "user,email,id"

## Trimming and Cleaning

Remove leading or trailing characters from a string.

	userInput := "  some important value\t\n"

	// Trim leading and trailing whitespace.
	cleaned := strings.TrimSpace(userInput)
	// cleaned is "some important value"

	// Trim a specific prefix or suffix.
	addr := "[INFO] This is a log message."
	msg := strings.TrimPrefix(addr, "[INFO] ")
	// msg is "This is a log message."

## Case Conversion

Change the case of a string.

	mixedCase := "Hello World"

	// Convert to upper or lower case.
	upper := strings.ToUpper(mixedCase) // "HELLO WORLD"
	lower := strings.ToLower(mixedCase) // "hello world"
*/
package strings
