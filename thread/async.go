package thread

import (
	"context"
	"errors"
)

// Async runs the provided function in a separate goroutine and returns a channel that will receive the result of the function.
// The function fn will be executed asynchronously and its result will be sent to the channel.
// The channel will be closed after the result is sent.
func Async[T any](f func() T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		ch <- f()
	}()
	return ch
}

// AsyncContext runs the provided function in a separate goroutine and returns a channel that will receive the result of the function.
// The function fn will be executed asynchronously and its result will be sent to the channel.
// The channel will be closed after the result is sent.
// If the context is done, the function will not be executed.
// The function will not be executed if the context is done.
func AsyncContext[T any](ctx context.Context, f func() T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		select {
		case ch <- f():
		case <-ctx.Done():
		}
	}()
	return ch

}

// AsyncOrErr runs the provided function in a separate goroutine and returns two channels: one that will receive the result of the function, and another that will receive any error that occurs during the function execution.
// The function fn will be executed asynchronously and its result and error will be sent to the respective channels.
// The channels will be closed after the result or error is sent.
// If the provided function is nil, an error will be sent to the error channel.
func AsyncOrErr[T any](f func() (T, error)) (<-chan T, <-chan error) {
	ch := make(chan T, 1)      // Channel to receive the result of the function
	err := make(chan error, 1) // Channel to receive any error that occurs during the function execution

	go func() {
		defer close(ch)  // Close the result channel when the goroutine finishes
		defer close(err) // Close the error channel when the goroutine finishes

		if f == nil { // Check if the provided function is nil
			err <- errors.New("function is nil") // Send an error to the error channel if the function is nil
			return
		}

		v, e := f()   // Call the provided function
		if e != nil { // Check if an error occurred during the function execution
			err <- e // Send the error to the error channel
			return
		}

		ch <- v // Send the result of the function to the result channel
	}()

	return ch, err // Return the result and error channels
}

// AsyncContextOrErr runs the provided function in a separate goroutine and returns two channels:
//
//	one that will receive the result of the function,
//	and another that will receive any error that occurs during the function execution.
//
// The function fn will be executed asynchronously and its result and error will be sent to the respective channels.
// The channels will be closed after the result or error is sent.
// If the provided function is nil, an error will be sent to the error channel.
// If the context is done, the function will not be executed.
// The function will not be executed if the context is done.
func AsyncContextOrErr[T any](ctx context.Context, f func(context.Context) (T, error)) (<-chan T, <-chan error) {
	ch := make(chan T, 1)      // Channel to receive the result of the function
	err := make(chan error, 1) // Channel to receive any error that occurs during the function execution

	go func() {
		defer close(ch)  // Close the result channel when the goroutine finishes
		defer close(err) // Close the error channel when the goroutine finishes
		if f == nil {    // Check if the provided function is nil
			err <- errors.New("function is nil") // Send an error to the error channel if the function is nil
			return
		}
		select {
		case <-ctx.Done():
			err <- ctx.Err()
			return
		default:
			v, e := f(ctx) // Call the provided function
			if e != nil {  // Check if an error occurred during the function execution
				err <- e // Send the error to the error channel
				return
			}
			ch <- v
		}
	}()
	return ch, err // Return the result and error channels
}
