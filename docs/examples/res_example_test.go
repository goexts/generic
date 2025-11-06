// Package examples demonstrates advanced Result-style helpers.
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

	// Example 4: Using Unpack
	r4 := res.Ok(42)
	v4, err4 := r4.Unpack()
	fmt.Println("v4:", v4, "err4:", err4)

	// Example 5: Using Of
	r5 := res.Of(42, nil)
	v5, err5 := r5.Unpack()
	fmt.Println("v5:", v5, "err5:", err5)
}

func ExampleResultHelpers() {
	ResultHelpers()

	// Output:
	// v1: 42
	// Error: boom
	// v3: 100
	// v4: 42 err4: <nil>
	// v5: 42 err5: <nil>
}

// Helper function to reduce boilerplate for handling results.
func HandleResult[T any](r res.Result[T]) {
	if val, ok := r.Ok(); ok {
		fmt.Println("Success:", val)
	} else {
		fmt.Println("Error:", r.Err())
	}
}

func ExampleHandleResult() {
	// Example 1: Using the helper function
	r1 := res.Ok(42)
	HandleResult(r1)

	r2 := res.Err[int](errors.New("boom"))
	HandleResult(r2)

	// Output:
	// Success: 42
	// Error: boom
}
