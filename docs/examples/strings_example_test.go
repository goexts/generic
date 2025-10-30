// Package examples demonstrates strings helpers.
package examples

import (
	"fmt"
	gostr "github.com/goexts/generic/strings"
)

func ExampleParseOr() {
	result := gostr.ParseOr("default", "")
	fmt.Println("ParseOr:", result)

	// Output:
	// ParseOr: default
}
