// Package examples demonstrates a simple Promise usage.
package examples

import (
	"errors"
	"fmt"
	"time"

	"github.com/goexts/generic/promise"
)

func ExampleNew() {
	p := promise.New(func(resolve func(string), reject func(error)) {
		go func() {
			time.Sleep(100 * time.Millisecond)
			resolve("done")
		}()
	})
	res, err := p.Await()
	fmt.Println("res:", res, "err:", err)

	p2 := promise.New(func(resolve func(int), reject func(error)) { reject(errors.New("fail")) })
	_, err2 := p2.Await()
	fmt.Println("err2 != nil:", err2 != nil)

	// Output:
	// res: done err: <nil>
	// err2 != nil: true
}
