// Package main demonstrates slices/bytes helpers.
package main

import (
	"fmt"

	"github.com/goexts/generic/slices/bytes"
)

func main() {
	b := []byte("hello")
	fmt.Println("contains 'e':", bytes.Contains(b, 'e'))
	b2 := bytes.Clone(b)
	fmt.Println("clone len:", len(b2))
}
