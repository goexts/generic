// Package main shows cond.If single-line conditional usage (Report Q4).
package main

import (
	"errors"
	"fmt"

	"github.com/goexts/generic/cond"
)

func GetStatusMessage(err error) string { // Q4
	return cond.If(err == nil, "Status: OK", "Status: Failed")
}

func main() {
	fmt.Println(GetStatusMessage(nil))
	fmt.Println(GetStatusMessage(errors.New("boom")))
}
