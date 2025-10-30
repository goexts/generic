// Package examples demonstrates slices/runes helpers.
package examples

import (
	"fmt"

	"github.com/goexts/generic/slices/runes"
)

func ExampleRunes() {
	r := runes.Runes("héllo")
	fmt.Println("index of 'é':", r.Index([]rune("é")))
	r2 := r.Clone()
	fmt.Println("clone size:", len(r2))

	// Output:
	// index of 'é': 1
	// clone size: 5
}
