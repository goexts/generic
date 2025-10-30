// Package main shows usage of ptr helpers.
package main

import (
	"fmt"

	"github.com/goexts/generic/ptr"
)

func main() {
	p := ptr.Of(42)
	fmt.Println("value via ptr:", *p)
	zero := ptr.Zero[int]()
	fmt.Println("zero:", *zero)
}
