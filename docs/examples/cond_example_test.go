// Package examples shows cond.If single-line conditional usage (Report Q4).
package examples

import (
	"errors"
	"fmt"

	"github.com/goexts/generic/cond"
)

// GetStatusMessage returns a status message based on the error.
// If err is nil, it returns "Status: OK", otherwise "Status: Failed".
// This demonstrates the use of cond.If for simple conditional expressions.
func GetStatusMessage(err error) string {
	return cond.If(err == nil, "Status: OK", "Status: Failed")
}

func ExampleGetStatusMessage() {
	fmt.Println(GetStatusMessage(nil))
	fmt.Println(GetStatusMessage(errors.New("something went wrong")))

	// Output:
	// Status: OK
	// Status: Failed
}

// If demonstrates the use of cond.If for simple conditional expressions.
func ExampleIf() {
	fmt.Println(cond.If(true, "Status: OK", "Status: Failed"))
	fmt.Println(cond.If(false, "Status: OK", "Status: Failed"))

	// Output:
	// Status: OK
	// Status: Failed
}
