/*
Package ptr provides generic utility functions for working with pointers.
It simplifies common operations such as creating a pointer from a literal value,
safely dereferencing a potentially nil pointer, and comparing pointer values.

# Creating Pointers with `Of`

In Go, you cannot take the address of a literal value directly. This package
solves that problem, which is especially useful for populating struct fields that
are pointers.

	type Config struct {
		Timeout *int
		Name    *string
	}

	// Verbose way without the ptr package:
	/*
	timeout := 30
	name := "default-name"
	cfg := Config{
		Timeout: &timeout,
		Name:    &name,
	}
	*/

	// Concise way with the ptr package:
	cfg := Config{
		Timeout: ptr.Of(30),
		Name:    ptr.Of("default-name"),
	}

# Safely Dereferencing with `Value`

Dereferencing a nil pointer causes a panic. The `Value` function provides a safe
way to get the value of a pointer, returning the zero value of the type if the
pointer is nil.

	var timeout *int // nil
	var name = ptr.Of("my-app")

	// Safely get the value or the zero value (0 for int).
	timeoutValue := ptr.Value(timeout) // Returns 0

	// Safely get the value of a non-nil pointer.
	nameValue := ptr.Value(name) // Returns "my-app"

# Comparing Pointers with `Equal`

The `Equal` function safely compares the values that two pointers point to.
It handles nil pointers gracefully.

	p1 := ptr.Of(100)
	p2 := ptr.Of(100)
	p3 := ptr.Of(200)
	var p4 *int // nil

	ptr.Equal(p1, p2) // true
	ptr.Equal(p1, p3) // false
	ptr.Equal(p1, p4) // false
	ptr.Equal(p4, p4) // true
*/
package ptr
