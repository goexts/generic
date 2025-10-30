// Package examples demonstrates all methods of the ptr package.
package examples

import (
	"fmt"

	"github.com/goexts/generic/ptr"
)

func PtrHelpers() {
	// Example 1: ptr.Of
	p := ptr.Of(42)
	fmt.Println("ptr.Of(42):", *p)

	// Example 2: ptr.Val
	val := ptr.Val(p)
	fmt.Println("ptr.Val(p):", val)

	// Example 3: ptr.To
	var anyValue any = "hello"
	strPtr := ptr.To[string](anyValue)
	fmt.Println("ptr.To[string](anyValue):", *strPtr)

	// Example 4: ptr.ToVal
	anyPtr := ptr.Of("world")
	strVal := ptr.ToVal[string](anyPtr)
	fmt.Println("ptr.ToVal[string](anyPtr):", strVal)

	// Example 5: Handling nil pointers
	nilPtr := (*int)(nil)
	nilVal := ptr.Val(nilPtr)
	fmt.Println("ptr.Val(nilPtr):", nilVal)
}

func ExamplePtrHelpers() {
	PtrHelpers()

	// Output:
	// ptr.Of(42): 42
	// ptr.Val(p): 42
	// ptr.To[string](anyValue): hello
	// ptr.ToVal[string](anyPtr): world
	// ptr.Val(nilPtr): 0
}
