/*
Package res provides a generic, Rust-inspired `Result[T]` type for expressive
error handling.

# Core Concept

A `Result[T]` is a type that represents either a success (containing a value
of type T) or a failure (containing an error). It is a monadic type that allows
for chaining operations in a clean, readable way, especially in data processing
pipelines where each step can fail.

This pattern provides an alternative to returning `(value, error)` pairs at each
step. Instead of checking for an error after every call, you can chain methods
and handle the final result once.

# Warning: Paradigm and Trade-offs

While powerful, the `Result` type introduces a paradigm that is not idiomatic
Go. Standard Go error handling (returning `(T, error)`) is simpler and more
direct for most use cases. The `Result` type is best suited for specific
scenarios like complex data transformation chains.

Be especially cautious with methods like `Unwrap` and `Expect`, which panic on
an `Err` value. They should only be used when an error is considered a fatal,
unrecoverable bug, similar to the `must` package.

# Example

	// A function that returns a Result
	func ParseInt(s string) res.Result[int] {
		n, err := strconv.Atoi(s)
		return res.Of(n, err)
	}

	// Chaining operations
	result := ParseInt("123").Unwrap() // result is 123

	// Safely handling the result
	val, ok := ParseInt("not-a-number").Ok()
	// val is 0, ok is false
*/
package res
