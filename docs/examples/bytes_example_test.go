// Package examples demonstrates slices/bytes helpers.
package examples

import (
	"fmt"

	"github.com/goexts/generic/slices/bytes"
)

func ExampleBytes() {
	b := bytes.Bytes("hello")
	fmt.Println("contains 'e':", b.Contains([]byte("e")))
	b2 := b.Clone()
	fmt.Println("clone len:", len(b2))

	// Output:
	// contains 'e': true
	// clone len: 5
}
