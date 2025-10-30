// Package examples demonstrates a simple Promise usage.
package examples

import (
	"errors"
	"fmt"
	"time"

	"github.com/goexts/generic/promise"
)

func ExampleNew() {
	// Create a new promise. The executor function is run in a new goroutine.
	p := promise.New(func(resolve func(string), _ func(error)) {
		// Simulate an asynchronous operation like a network call.
		time.Sleep(100 * time.Millisecond)
		resolve("done")
	})

	// Await blocks until the promise is settled (fulfilled or rejected).
	res, err := p.Await()
	fmt.Println("res:", res, "err:", err)

	// Create a promise that immediately rejects.
	p2 := promise.New(func(_ func(int), reject func(error)) {
		reject(errors.New("fail"))
	})
	_, err2 := p2.Await()
	fmt.Println("err2 != nil:", err2 != nil)

	// Output:
	// res: done err: <nil>
	// err2 != nil: true
}
