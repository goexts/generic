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
