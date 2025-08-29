package promise

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestPromise_AsyncAwait(t *testing.T) {
	t.Run("Successful execution", func(t *testing.T) {
		p := Async(func() (int, error) {
			time.Sleep(10 * time.Millisecond)
			return 42, nil
		})

		val, err := Await(p)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if val != 42 {
			t.Errorf("Expected value 42, but got %v", val)
		}
	})

	t.Run("Execution with error", func(t *testing.T) {
		expectedErr := errors.New("something went wrong")
		p := Async(func() (int, error) {
			return 0, expectedErr
		})

		_, err := Await(p)

		if !errors.Is(err, expectedErr) {
			t.Errorf("Expected error '%v', but got '%v'", expectedErr, err)
		}
	})
}

func TestPromise_Then(t *testing.T) {
	t.Run("Chain of Then calls", func(t *testing.T) {
		p := Async(func() (int, error) {
			return 10, nil
		}).Then(func(val int) int {
			return val * 2
		}).Then(func(val int) int {
			return val + 5
		})

		finalVal, err := Await(p)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if finalVal != 25 { // 10*2 + 5
			t.Errorf("Expected final value 25, but got %v", finalVal)
		}
	})
}

func TestPromise_Catch(t *testing.T) {
	t.Run("Catch and recover from an error", func(t *testing.T) {
		initialErr := errors.New("initial error")
		p := Async(func() (int, error) {
			return 0, initialErr
		}).Catch(func(err error) (int, error) {
			if errors.Is(err, initialErr) {
				// Recover with a new value
				return 100, nil
			}
			// Propagate other errors
			return 0, err
		})

		val, err := Await(p)

		if err != nil {
			t.Errorf("Expected no error after catch, but got %v", err)
		}
		if val != 100 {
			t.Errorf("Expected recovered value 100, but got %v", val)
		}
	})

	t.Run("Catch propagates new error", func(t *testing.T) {
		newErr := errors.New("new error")
		p := Async(func() (int, error) {
			return 0, errors.New("original error")
		}).Catch(func(_ error) (int, error) {
			// Return a new error
			return 0, newErr
		})

		_, err := Await(p)

		if !errors.Is(err, newErr) {
			t.Errorf("Expected new error, but got %v", err)
		}
	})
}

func TestPromise_PanicSafety(t *testing.T) {
	t.Run("Executor panics", func(t *testing.T) {
		p := New(func(_ func(int), _ func(error)) {
			panic("executor panic")
		})

		_, err := Await(p)

		if err == nil {
			t.Error("Expected an error from panic, but got nil")
		} else if !strings.Contains(err.Error(), "promise executor panicked") {
			t.Errorf("Error message should indicate a panic, but got: %v", err)
		}
	})
}

func TestPromise_ThenWithPromise(t *testing.T) {
	t.Run("Chaining with another promise", func(t *testing.T) {
		p := Async(func() (int, error) {
			return 2, nil
		}).ThenWithPromise(func(val int) *Promise[int] {
			return Async(func() (int, error) {
				time.Sleep(10 * time.Millisecond)
				return val * val, nil // 2*2 = 4
			})
		})

		finalVal, err := Await(p)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if finalVal != 4 {
			t.Errorf("Expected final value 4, but got %v", finalVal)
		}
	})
}
