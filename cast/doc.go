/*
Package cast provides safe, generic alternatives to Go's standard type assertion.

This package simplifies type assertions by providing convenient, single-expression
functions that handle the `value, ok` idiom in different ways, such as returning
a default value or a zero value upon failure.

Example:

	var myVal any = "hello"

	// Standard Go type assertion
	str, ok := myVal.(string)
	if !ok {
		str = "default"
	}

	// With cast package - cleaner and more concise
	str = cast.Or(myVal, "default")
*/
package cast
