// Package examples shows cond.If single-line conditional usage (Report Q4).
package examples

import (
	"fmt"

	"github.com/goexts/generic/cond"
)

func ExampleIf() {
	fmt.Println(cond.If(true, "Status: OK", "Status: Failed"))
	fmt.Println(cond.If(false, "Status: OK", "Status: Failed"))

	// Output:
	// Status: OK
	// Status: Failed
}
