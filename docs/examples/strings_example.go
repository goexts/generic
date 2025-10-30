// Package main demonstrates strings helpers.
package main

import (
	"fmt"

	gostr "github.com/goexts/generic/strings"
)

func main() {
	v, ok := gostr.Atoi("42")
	fmt.Println("Atoi:", v, ok)
	_, bad := gostr.Atoi("x")
	fmt.Println("Atoi bad ok=false:", bad)
}
