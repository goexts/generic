/*
Package must provides helper functions that wrap calls returning an error and
panic if the error is non-nil. This is intended to reduce boilerplate in specific,
controlled contexts.

WARNING: The functions in this package should be used with extreme care. They
intentionally convert a recoverable error into a non-recoverable panic. This is
only appropriate in specific situations where an error is considered a bug in
the program, not a predictable runtime failure.

Appropriate Use Cases:

  - Program initialization (e.g., in `init` functions or `main`):
    parsing hardcoded configuration, compiling essential regular expressions,
    or setting up database connections that are required for the application
    to start.

  - Test setup: When setting up test fixtures where a failure indicates a
    broken test, not a feature to be tested.

Example:

	// Instead of:
	// re, err := regexp.Compile(`\w+`)
	// if err != nil {
	// 	panic(err)
	// }

	// Use must.Must for cleaner initialization code:
	re := must.Must(regexp.Compile(`\w+`))

DO NOT use these functions for regular application logic where errors are
expected (e.g., handling user input, network requests, file I/O).
*/
package must
