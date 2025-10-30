// Package main demonstrates simple Set operations.
package main

import (
	"fmt"

	"github.com/goexts/generic/set"
)

func main() {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)
	fmt.Println("has 1:", s.Has(1))
	s.Remove(1)
	fmt.Println("size:", s.Len())
}
