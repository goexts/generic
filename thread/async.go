package thread

// Async runs the provided function in a separate goroutine and returns a channel that will receive the result of the function.
// The function f will be executed asynchronously and its result will be sent to the channel.
// The channel will be closed after the result is sent.
func Async[T any](f func() T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		ch <- f()
	}()
	return ch
}
