// Package examples demonstrates Promise usage with advanced scenarios.
package examples

import (
	"errors"
	"fmt"
	"time"

	"github.com/goexts/generic/promise"
	"github.com/goexts/generic/res"
)

func ExampleNew() {
	// Basic promise with resolve and reject.
	p := promise.New(func(resolve func(string), _ func(error)) {
		time.Sleep(100 * time.Millisecond)
		resolve("done")
	})
	res1, err := p.Await()
	fmt.Println("res1:", res1, "err:", err)

	// Promise with immediate rejection.
	p2 := promise.New(func(_ func(int), reject func(error)) {
		reject(errors.New("fail"))
	})
	_, err2 := p2.Await()
	fmt.Println("err2 != nil:", err2 != nil)

	// Output:
	// res1: done err: <nil>
	// err2 != nil: true
}

func ExampleChain() {
	// Chain promises for sequential operations.
	p1 := promise.New(func(resolve func(int), _ func(error)) {
		time.Sleep(100 * time.Millisecond)
		resolve(10)
	})

	p2 := promise.Then(p1, func(val int) (string, error) {
		return fmt.Sprintf("value: %d", val), nil
	})

	res2, err := p2.Await()
	fmt.Println("res2:", res2, "err:", err)

	// Output:
	// res2: value: 10 err: <nil>
}

func PromiseAllHelpers() {
	// Run promises concurrently and wait for all.
	p1 := promise.New(func(resolve func(int), _ func(error)) {
		time.Sleep(200 * time.Millisecond)
		resolve(1)
	})

	p2 := promise.New(func(resolve func(int), _ func(error)) {
		time.Sleep(100 * time.Millisecond)
		resolve(2)
	})

	results := promise.All(p1, p2)
	ints, err := results.Await()
	fmt.Println("ints:", ints, "err:", err)
}

func ExamplePromiseAllHelpers() {
	PromiseAllHelpers()

	// Output:
	// ints: [1 2] err: <nil>
}

func UnwrapOrHelpers() {
	p := promise.New(func(resolve func(res.Result[int]), _ func(error)) {
		time.Sleep(100 * time.Millisecond)
		resolve(res.Ok(42))
	})

	result, err := p.Await()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	val := result.UnwrapOr(0)
	fmt.Println("val:", val)
}

func ExampleUnwrapOrHelpers() {
	// Combine promise with res.Result for robust error handling.
	UnwrapOrHelpers()

	// Output:
	// val: 42
}
