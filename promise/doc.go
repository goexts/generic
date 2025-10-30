/*
Package promise provides a generic, type-safe implementation of Promises, inspired
by similar concepts in other languages like JavaScript. It is designed to simplify
the management of asynchronous operations and avoid complex callback chains
(often referred to as "callback hell").

# Core Concepts

A Promise represents the eventual completion (or failure) of an asynchronous
operation and its resulting value. A Promise is always in one of three states:

- **pending**: The initial state; the asynchronous operation has not yet completed.
- **fulfilled**: The operation completed successfully, and the promise now has a result value.
- **rejected**: The operation failed, and the promise now holds an error.

# Basic Usage

The primary way to create a promise is with the `New` function, which takes an
`executor` function. The executor is run in a new goroutine and is given `resolve`
and `reject` functions to control the promise's final state.

	// Create a promise that resolves with a message after a short delay.
	p := promise.New(func(resolve func(string), reject func(error)) {
		time.Sleep(100 * time.Millisecond)
		resolve("Hello from the promise!")
	})

	// The Await method blocks until the promise is settled (either fulfilled or rejected).
	val, err := p.Await()
	// val will be "Hello from the promise!", err will be nil.

# Chaining Asynchronous Operations

Promises shine when you need to orchestrate a sequence of asynchronous steps.
Methods like `Then`, `Catch`, and `Finally` allow you to build a clean, linear
workflow.

Example of a complete chain:

	// 1. Start an async operation to fetch a user ID.
	userIDPromise := promise.Async(func() (int, error) {
		fmt.Println("Fetching user ID...")
		time.Sleep(50 * time.Millisecond)
		return 123, nil // Simulate success
	})

	// 2. Use `Then` to fetch user data once the ID is available.
	userDataPromise := promise.Then(userIDPromise, func(id int) (string, error) {
		fmt.Printf("Fetching data for user %d...\n", id)
		time.Sleep(50 * time.Millisecond)
		return fmt.Sprintf("{\"name\": \"Alice\", \"id\": %d}", id), nil
	})

	// 3. Use `Catch` to handle any errors that might have occurred in the chain.
	finalPromise := promise.Catch(userDataPromise, func(err error) (string, error) {
		fmt.Printf("An error occurred: %v. Recovering.\n", err)
		return "default data", nil // Recover with a default value
	})

	// 4. Await the final result.
	finalResult, _ := finalPromise.Await()
	fmt.Printf("Final result: %s\n", finalResult)

This creates a readable, non-blocking sequence of dependent operations.
*/
package promise
