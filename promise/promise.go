package promise

import (
	"fmt"
	"sync"
)

// Promise represents the eventual completion (or failure) of an asynchronous
// operation and its resulting value. It is a generic, type-safe implementation
// inspired by the JavaScript Promise API.
type Promise[T any] struct {
	lock    sync.Mutex
	value   T
	err     error
	done    chan struct{}
	settled bool // To prevent multiple resolves/rejects
}

// New creates a new Promise. The provided executor function is executed in a new
// goroutine. The executor receives `resolve` and `reject` functions to control
// the promise's outcome.
func New[T any](executor func(resolve func(T), reject func(error))) *Promise[T] {
	p := &Promise[T]{
		done: make(chan struct{}),
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// Automatically reject if the executor panics.
				p.Reject(fmt.Errorf("promise executor panicked: %v", r))
			}
		}()
		executor(p.Resolve, p.Reject)
	}()

	return p
}

// Resolve fulfills the promise with a value. If the promise is already settled,
// this call is ignored.
func (p *Promise[T]) Resolve(value T) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.settled {
		return
	}

	p.value = value
	p.settled = true
	close(p.done)
}

// Reject rejects the promise with an error. If the promise is already settled,
// this call is ignored.
func (p *Promise[T]) Reject(err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.settled {
		return
	}

	p.err = err
	p.settled = true
	close(p.done)
}

// Await blocks until the promise is settled and returns the resulting value and
// error. It is the primary way to get the result of a promise.
func (p *Promise[T]) Await() (T, error) {
	<-p.done
	return p.value, p.err
}

// Then attaches a callback that executes when the promise is fulfilled.
// It returns a new promise that resolves with the result of the onFulfilled callback.
// If the original promise is rejected, the new promise is rejected with the same error.
func (p *Promise[T]) Then(onFulfilled func(T) T) *Promise[T] {
	return New(func(resolve func(T), reject func(error)) {
		val, err := p.Await()
		if err != nil {
			reject(err)
			return
		}
		defer func() {
			if r := recover(); r != nil {
				reject(fmt.Errorf("panic in Then: %v", r))
			}
		}()
		resolve(onFulfilled(val))
	})
}

// ThenWithPromise is like Then, but the callback returns a new Promise.
// This allows for chaining of asynchronous operations.
func (p *Promise[T]) ThenWithPromise(onFulfilled func(T) *Promise[T]) *Promise[T] {
	return New(func(resolve func(T), reject func(error)) {
		val, err := p.Await()
		if err != nil {
			reject(err)
			return
		}
		defer func() {
			if r := recover(); r != nil {
				reject(fmt.Errorf("panic in ThenWithPromise: %v", r))
			}
		}()
		// Chain the promise
		newPromise := onFulfilled(val)
		newVal, newErr := newPromise.Await()
		if newErr != nil {
			reject(newErr)
		} else {
			resolve(newVal)
		}
	})
}

// Catch attaches a callback that executes when the promise is rejected.
// It allows for error handling and recovery. The onRejected callback can return a
// new value to fulfill the promise, or a new error to continue the rejection chain.
func (p *Promise[T]) Catch(onRejected func(error) (T, error)) *Promise[T] {
	return New(func(resolve func(T), reject func(error)) {
		val, err := p.Await()
		if err != nil {
			defer func() {
				if r := recover(); r != nil {
					reject(fmt.Errorf("panic in Catch: %v", r))
				}
			}()
			newVal, newErr := onRejected(err)
			if newErr != nil {
				reject(newErr)
			} else {
				resolve(newVal)
			}
			return
		}
		resolve(val)
	})
}

// Finally attaches a callback that executes when the promise is settled (either
// fulfilled or rejected). It is useful for cleanup logic.
// The returned promise will be settled with the same value or error as the
// original promise, after onFinally has completed.
func (p *Promise[T]) Finally(onFinally func()) *Promise[T] {
	return New(func(resolve func(T), reject func(error)) {
		val, err := p.Await()

		var panicErr error
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicErr = fmt.Errorf("panic in Finally: %v", r)
				}
			}()
			onFinally()
		}()

		if panicErr != nil {
			reject(panicErr)
			return
		}

		if err != nil {
			reject(err)
		} else {
			resolve(val)
		}
	})
}

// Async is a helper function that wraps a function returning (T, error)
// into a new Promise. This is useful for converting existing functions into
// promise-based asynchronous operations.
func Async[T any](f func() (T, error)) *Promise[T] {
	return New(func(resolve func(T), reject func(error)) {
		val, err := f()
		if err != nil {
			reject(err)
		} else {
			resolve(val)
		}
	})
}

// Await is a standalone function that waits for a promise to be settled.
// It is a functional equivalent of the p.Await() method.
func Await[T any](p *Promise[T]) (T, error) {
	return p.Await()
}
