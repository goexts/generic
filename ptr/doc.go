/*
Package ptr provides generic utility functions for working with pointers.
It simplifies common operations such as creating a pointer from a literal value
or safely dereferencing a pointer that might be nil.

This package is particularly useful in contexts where you need to assign a
pointer to a struct field, but you only have a literal value (e.g., a string,
an integer, or a boolean).

Example:

	// Without ptr package
	port := 8080
	config := &ServerConfig{
		Port: &port,
	}

	// With ptr package
	config := &ServerConfig{
		Port: ptr.Of(8080), // Cleaner and more concise
	}
*/
package ptr
