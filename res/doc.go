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
scenarios like complex data transformation chains where the flow of data is
the primary concern.

Be especially cautious with methods like `Unwrap` and `Expect`, which panic on
an `Err` value. They should only be used when an error is considered a fatal,
unrecoverable bug, similar to the `must` package.

# Usage: Data Processing Pipeline

The primary benefit of `Result` is in chaining operations where any step can fail.
The chain short-circuits as soon as an error occurs.

Consider a sequence of operations:
1. Get a filename from a map.
2. Read the file content.
3. Parse the content into a number.

	// Define helper functions that each return a Result.
	func getFilename(config map[string]string) res.Result[string] {
	    if name, ok := config["filename"]; ok {
	        return res.Ok(name)
	    }
	    return res.Err[string](errors.New("filename not found in config"))
	}

	func readContent(filename string) res.Result[string] {
	    // Simulate reading a file.
	    if filename == "data.txt" {
	        return res.Ok("12345")
	    }
	    return res.Err[string](fmt.Errorf("file not found: %s", filename))
	}

	func parseNumber(content string) res.Result[int] {
	    n, err := strconv.Atoi(content)
	    return res.Of(n, err)
	}

	// Now, chain these operations together.
	config := map[string]string{"filename": "data.txt"}

	// The `AndThen` method chains functions that return a Result.
	// The chain stops at the first `Err`.
	finalResult := getFilename(config).
	    AndThen(readContent).
	    AndThen(parseNumber)

	// Safely handle the outcome.
	if finalResult.IsErr() {
	    fmt.Printf("Pipeline failed: %v\n", finalResult.Err())
	} else {
	    // No error occurred, we can safely get the value.
	    fmt.Printf("Pipeline succeeded, result: %d\n", finalResult.Unwrap())
	}
*/
package res
