package must_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/must"
)

func TestDo(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			val := must.Do(10, nil)
			assert.Equal(t, 10, val)
		})
	})

	t.Run("with error", func(t *testing.T) {
		err := errors.New("test error")
		assert.PanicsWithValue(t, err, func() {
			must.Do(0, err)
		})
	})
}

func TestDo2(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			v1, v2 := must.Do2("hello", 123, nil)
			assert.Equal(t, "hello", v1)
			assert.Equal(t, 123, v2)
		})
	})

	t.Run("with error", func(t *testing.T) {
		err := errors.New("test error 2")
		assert.PanicsWithValue(t, err, func() {
			must.Do2("", 0, err)
		})
	})
}

func TestCast(t *testing.T) {
	var i any = "world"

	// Successful cast
	assert.NotPanics(t, func() {
		val := must.Cast[string](i)
		assert.Equal(t, "world", val)
	})

	// Failed cast
	assert.Panics(t, func() {
		must.Cast[int](i)
	})
}
