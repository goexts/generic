// Package main demonstrates Result-style helpers.
package main

import (
	"errors"
	"fmt"

	"github.com/goexts/generic/res"
)

func mightFail(ok bool) res.Result[int] {
	if ok {
		return res.Ok(1)
	}
	return res.Err[int](errors.New("boom"))
}

func main() {
	r1 := mightFail(true)
	v1, e1 := r1.Unwrap()
	fmt.Println("v1:", v1, e1 == nil)
	r2 := mightFail(false)
	v2, e2 := r2.Unwrap()
	fmt.Println("v2:", v2, e2 != nil)
}
