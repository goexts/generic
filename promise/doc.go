/*
Package promise provides a generic, type-safe implementation of Promises, inspired
by the JavaScript Promise API. It is designed to simplify the management of
asynchronous operations and avoid complex callback chains ("callback hell").

# Core Concepts

A Promise represents the eventual completion (or failure) of an asynchronous
operation and its resulting value. A Promise is in one of three states:

  - pending: the initial state; neither fulfilled nor rejected.
  - fulfilled: the operation completed successfully, resulting in a value.
  - rejected: the operation failed, resulting in an error.

# Basic Usage

The primary way to create a promise is with the `New` function, which takes
an `executor` function. The executor is run in a new goroutine and receives
`resolve` and `reject` functions to control the promise's outcome.

Example:

	// Create a promise that resolves after a delay.
	p := promise.New(func(resolve func(string), reject func(error)) {
		time.Sleep(100 * time.Millisecond)
		resolve("Hello, World!")
	})

	// The Await method blocks until the promise is settled.
	val, err := p.Await()
	// val is "Hello, World!", err is nil.

# Chaining

Promises can be chained together using methods like `Then`, `Catch`, and `Finally`
to create a clean, linear asynchronous workflow.

	resultPromise := promise.Async(func() (int, error) {
		// Simulate an API call
		return 42, nil
	}).Then(func(val int) int {
		// Transform the result
		return val * 2
	}).Catch(func(err error) (int, error) {
		// Handle any previous errors and potentially recover.
		fmt.Printf("An error occurred: %v\n", err)
		return 0, nil // Recover with a default value
	})

	finalResult, _ := resultPromise.Await() // finalResult is 84
*/
package promise
