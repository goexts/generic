package promise

import (
	"fmt"
	"sync"
)

// Promise represents the eventual completion (or failure) of an asynchronous operation and its resulting value.
type Promise[T any] struct {
	lock    sync.Mutex
	value   T
	err     error
	done    chan struct{}
	settled bool // To prevent multiple resolves/rejects
}

// New creates a new Promise. The provided executor function is executed in a new goroutine.
// The executor receives resolve and reject functions to control the promise's outcome.
func New[T any](executor func(resolve func(T), reject func(error))) *Promise[T] {
	p := &Promise[T]{
		done: make(chan struct{}),
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// automatically reject if the executor panics
				p.Reject(fmt.Errorf("promise executor panicked: %v", r))
			}
		}()
		executor(p.Resolve, p.Reject)
	}()

	return p
}

// Resolve fulfills the promise with a value.
// It's safe to call multiple times, but only the first call will have an effect.
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

// Reject rejects the promise with an error.
// It's safe to call multiple times, but only the first call will have an effect.
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

// Await waits for the promise to be settled and returns the value and error.
// This is the equivalent of `await` in other languages.
func (p *Promise[T]) Await() (T, error) {
	<-p.done
	return p.value, p.err
}

// Then attaches a callback for the resolution of the Promise.
// It returns a new Promise resolving with the result of the callback.
// If the current promise is rejected, the new promise is also rejected with the same error.
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

// Catch handles promise rejections.
// It returns a new Promise. If the original promise is fulfilled, the new promise is fulfilled with the same value.
// If the original promise is rejected, the onRejected handler is called, which can recover from the error.
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

// Async is a helper to create and run a promise for a function that returns a value and an error.
// This is your `async` function.
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

// Await is a standalone function to wait for a promise.
func Await[T any](p *Promise[T]) (T, error) {
	return p.Await()
}
