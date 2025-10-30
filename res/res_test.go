package res_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/res"
)

func TestResult(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		r := res.Ok(42)
		assert.True(t, r.IsOk())
		assert.False(t, r.IsErr())
		assert.Equal(t, 42, r.Unwrap())
		assert.Nil(t, r.Err())

		val, ok := r.Ok()
		assert.True(t, ok)
		assert.Equal(t, 42, val)
	})

	t.Run("Err", func(t *testing.T) {
		err := errors.New("something went wrong")
		r := res.Err[int](err)

		assert.False(t, r.IsOk())
		assert.True(t, r.IsErr())
		assert.Equal(t, err, r.Err())

		val, ok := r.Ok()
		assert.False(t, ok)
		assert.Equal(t, 0, val) // Should be the zero value for int
	})

	t.Run("Of", func(t *testing.T) {
		// Success case
		r1 := res.Of(100, nil)
		assert.True(t, r1.IsOk())
		assert.Equal(t, 100, r1.Unwrap())

		// Error case
		err := errors.New("failure")
		r2 := res.Of(0, err)
		assert.True(t, r2.IsErr())
		assert.Equal(t, err, r2.Err())
	})

	t.Run("Unwrap", func(t *testing.T) {
		assert.Equal(t, "hello", res.Ok("hello").Unwrap())
		assert.Panics(t, func() {
			res.Err[string](errors.New("panic error")).Unwrap()
		})
	})

	t.Run("UnwrapOr", func(t *testing.T) {
		assert.Equal(t, "world", res.Ok("world").UnwrapOr("default"))
		assert.Equal(t, "default", res.Err[string](errors.New("err")).UnwrapOr("default"))
	})

	t.Run("Expect", func(t *testing.T) {
		assert.Equal(t, 123, res.Ok(123).Expect("should not panic"))
		assert.PanicsWithValue(t, "custom message: an error occurred", func() {
			res.Err[int](errors.New("an error occurred")).Expect("custom message")
		})
	})

	t.Run("ExampleUsage", func(t *testing.T) {
		divide := func(a, b int) res.Result[int] {
			if b == 0 {
				return res.Err[int](errors.New("division by zero"))
			}
			return res.Ok(a / b)
		}

		// Good division
		result1 := divide(10, 2)
		assert.Equal(t, 5, result1.UnwrapOr(0))

		// Bad division
		result2 := divide(10, 0)
		assert.True(t, result2.IsErr())
		fmt.Println(result2.Err()) // Prints: division by zero
	})
}

func TestHelpers(t *testing.T) {
	t.Run("Or", func(t *testing.T) {
		// Success case
		val1 := res.Or("value", nil, "default")
		assert.Equal(t, "value", val1)

		// Error case
		val2 := res.Or("", errors.New("error"), "default")
		assert.Equal(t, "default", val2)
	})

	t.Run("OrZero", func(t *testing.T) {
		// Success case
		val1 := res.OrZero(123, nil)
		assert.Equal(t, 123, val1)

		// Error case
		val2 := res.OrZero(0, errors.New("error"))
		assert.Equal(t, 0, val2)
	})
}
