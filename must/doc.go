/*
Package must provides helper functions that wrap calls returning `(T, error)`
and panic if the error is non-nil. This is intended to reduce boilerplate code
in specific, controlled contexts where an error should never happen.

# Warning: Use with Extreme Care

The functions in this package intentionally convert a recoverable error into a
non-recoverable panic. This is an anti-pattern in normal Go application code.
It should only be used in situations where an error is truly unexpected and
indicates a critical, unrecoverable programmer error (e.g., a bug).

# Appropriate Use Cases

1.  **Program Initialization:** During startup (e.g., in `init` functions or at
    the top of `main`), when a failure means the application cannot run at all.

2.  **Test Setup:** When preparing test fixtures, where a failure indicates a
    broken test environment, not a feature to be tested.

## Example: Compiling a Regular Expression

It is common to compile regular expressions at the package level. Since the
pattern is hardcoded, a compilation failure is a programmer error, not a
runtime error. `must.Must` simplifies this.

	// Before: Verbose error handling for a panic-worthy error.
	/*
	var wordRegexp *regexp.Regexp

	func init() {
		var err error
		wordRegexp, err = regexp.Compile(`\w+`)
		if err != nil {
			panic(fmt.Sprintf("failed to compile word regexp: %v", err))
		}
	}
	*/

	// After: Using must.Must for concise, clear initialization.
	var wordRegexp = must.Must(regexp.Compile(`\w+`))

# Inappropriate Use Cases

NEVER use these functions for regular application logic where errors are
expected and should be handled gracefully. This includes, but is not limited to:

- Handling user input.
- Processing network requests or responses.
- Reading from or writing to files.
*/
package must
