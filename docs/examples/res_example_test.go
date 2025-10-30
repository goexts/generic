// Package examples demonstrates Result-style helpers.
package examples

import (
	"errors"
	"fmt"

	"github.com/goexts/generic/res"
)

func ResultHelpers() {
	// Example 1: Successful result
	r1 := res.Ok(42)
	if v1, ok := r1.Ok(); ok {
		fmt.Println("v1:", v1)
	}

	// Example 2: Error result - safe handling
	r2 := res.Err[int](errors.New("boom"))
	if _, ok := r2.Ok(); ok {
		fmt.Println("This won't be printed")
	} else {
		fmt.Println("Error:", r2.Err())
	}

	// Example 3: Using UnwrapOr
	r3 := res.Err[int](errors.New("oops"))
	v3 := r3.UnwrapOr(100) // Returns default value on error
	fmt.Println("v3:", v3)
}

func ExampleResultHelpers() {
	ResultHelpers()

	// Output:
	// v1: 42
	// Error: boom
	// v3: 100
}
