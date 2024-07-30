package thread

import (
	"context"
)

// Wait waits for a value to be available on the provided channel.
//
// The channel v must be a receive-only channel of type T. The function blocks
// until a value is received from the channel, and then returns that value.
//
// Parameters:
//   - v: A receive-only channel of type T.
//
// Returns:
//   - T: The value received from the channel.
func Wait[T any](v <-chan T) T {
	return <-v
}

// WaitWithContext waits for a value to be available on the provided channel
// within the given context.
//
// The channel v must be a receive-only channel of type T. The function blocks
// until a value is received from the channel or the context is done.
//
// Parameters:
//   - ctx: The context to control the waiting process.
//   - v: A receive-only channel of type T.
//
// Returns:
//   - T: The value received from the channel if successful.
//   - error: The error that caused the waiting process to be done.
func WaitWithContext[T any](ctx context.Context, v <-chan T) (T, error) {
	var ret T
	select {
	case ret = <-v: // If a value is received from the channel, return it.
		return ret, nil
	case <-ctx.Done(): // If the context is done, return an error.
		return ret, ctx.Err()
	}
}

// WaitOrErr waits for a value to be available on the provided channel or an
// error to be available on the provided error channel.
//
// The channel v must be a receive-only channel of type T. The function blocks
// until a value is received from the channel or an error is received from the
// error channel, and then returns the value or the error received from the
// channel.
//
// Parameters:
//   - v: A receive-only channel of type T.
//   - err: A receive-only channel of type error.
//
// Returns:
//   - T: The value received from the channel if successful.
//   - error: The error received from the error channel if successful.
func WaitOrErr[T any](v <-chan T, err <-chan error) (T, error) {
	var ret T
	select {
	case ret = <-v: // If a value is received from the channel, return it.
		return ret, nil
	case e := <-err: // If an error is received from the error channel, return it.
		return ret, e
	}
}

// WaitTimeoutOrErr waits for a value to be available on the provided channel within the given context.
// If a value is received from the channel, it is returned along with a nil error.
// If an error is received from the error channel, it is returned along with the received error.
// If the context is done, the function returns the last received value and an error indicating the context is done.
//
// Parameters:
//   - ctx: The context to control the waiting process.
//   - v: A receive-only channel of type T.
//   - err: A receive-only channel of type error.
//
// Returns:
//   - T: The value received from the channel if successful.
//   - error: The error that caused the waiting process to be done.
func WaitTimeoutOrErr[T any](ctx context.Context, v <-chan T, err <-chan error) (T, error) {
	var ret T
	select {
	case ret = <-v: // If a value is received from the channel, return it.
		return ret, nil
	case e := <-err: // If an error is received from the error channel, return it.
		return ret, e
	case <-ctx.Done(): // If the context is done, return an error.
		return ret, ctx.Err()
	}
}
