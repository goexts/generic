package cast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/cast"
)

func TestTry(t *testing.T) {
	var i any = 10

	// Successful cast
	val, ok := cast.Try[int](i)
	assert.True(t, ok)
	assert.Equal(t, 10, val)

	// Failed cast
	_, ok = cast.Try[string](i)
	assert.False(t, ok)
}

func TestOr(t *testing.T) {
	var i any = "hello"

	// Successful cast
	val1 := cast.Or[string](i, "default")
	assert.Equal(t, "hello", val1)

	// Failed cast
	val2 := cast.Or[int](i, 123)
	assert.Equal(t, 123, val2)
}

func TestOrZero(t *testing.T) {
	var i any = 10.5

	// Successful cast
	val1 := cast.OrZero[float64](i)
	assert.Equal(t, 10.5, val1)

	// Failed cast
	val2 := cast.OrZero[int](i)
	assert.Equal(t, 0, val2)
}
