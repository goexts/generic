// Package main demonstrates slices/runes helpers.
package main

import (
	"fmt"

	"github.com/goexts/generic/slices/runes"
)

func main() {
	r := runes.Runes("héllo")
	fmt.Println("index of 'é':", r.Index([]rune("é")))
	//r2 := r.Clone(r)
	//fmt.Println("clone size:", len(r2))
}
