// Package thread provides a simple way to run functions in goroutines with error handling.
package thread

import (
	"context"
)

// Caller is an interface that represents a function call that can be run in a goroutine.
// It allows for error handling and chaining of functions to be called after the function call completes.
type Caller[T any] interface {
	// Then sets a function to be called after the function call completes successfully.
	// The result of the function call is passed as an argument to the function.
	Then(func(T)) Caller[T]

	// Catch sets a function to be called if an error occurs during the function call.
	// The error is passed as an argument to the function.
	Catch(func(error)) Caller[T]

	// Finally sets a function to be called after the function call completes, regardless of whether an error occurred or not.
	Finally(func())
}

// caller is a struct that represents a function call to be run in a goroutine.
type caller[T any] struct {
	ctx     context.Context   // The context to use for the function call.
	f       func() (T, error) // The function to be called.
	then    []func(T)         // The function to be called after the function call completes successfully.
	catch   func(error)       // The function to be called if an error occurs during the function call.
	finally func()            // The function to be called after the function call completes, regardless of whether an error occurred or not.
}

// Then sets the function to be called after the function call completes successfully.
// The result of the function call is passed as an argument to the function.
func (c *caller[T]) Then(f func(T)) Caller[T] {
	c.then = append(c.then, f)
	return c
}

// Catch sets the function to be called if an error occurs during the function call.
// The error is passed as an argument to the function.
func (c *caller[T]) Catch(f func(error)) Caller[T] {
	c.catch = f
	return c
}

// Finally sets the function to be called after the function call completes, regardless of whether an error occurred or not.
func (c *caller[T]) Finally(f func()) {
	c.finally = f
}

// runGoroutine runs the function in a goroutine and handles error handling and chaining of functions.
func (c *caller[T]) runGoroutine() {
	go func() {
		defer func() {
			if c.finally != nil {
				c.finally()
			}
		}()
		ret, err := AsyncOrErr(func() (T, error) {
			return c.f()
		})

		t, e := WaitTimeoutOrErr(c.ctx, ret, err)
		if e != nil {
			if c.catch != nil {
				c.catch(e)
			}
			return
		}
		var then func(T)
		for _, then = range c.then {
			then(t)
		}
	}()
}

// Try creates a new caller with the given function and runs it in a goroutine.
// The caller can be used to chain functions to be called after the function call completes.
func Try[T any](f func() (T, error)) (Caller[T], func()) {
	c := &caller[T]{
		ctx: context.Background(),
		f:   f,
	}
	return c, c.runGoroutine
}

// TryWithContext creates a new caller with the given function and context and runs it in a goroutine.
// The caller can be used to chain functions to be called after the function call completes.
func TryWithContext[T any](ctx context.Context, f func() (T, error)) (Caller[T], func()) {
	c := &caller[T]{
		ctx: ctx,
		f:   f,
	}
	return c, c.runGoroutine
}
