package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/ptr"
)

func TestOf(t *testing.T) {
	val := 10
	p := ptr.Of(val)
	assert.NotNil(t, p)
	assert.Equal(t, val, *p)
}

func TestVal(t *testing.T) {
	val := "hello"
	p := &val
	assert.Equal(t, val, ptr.Val(p))

	var nilPtr *string
	assert.Equal(t, "", ptr.Val(nilPtr)) // Should return zero value for string
}

func TestTo(t *testing.T) {
	// Case 1: Input is of type T
	valT := 100
	p1 := ptr.To[int](valT)
	assert.NotNil(t, p1)
	assert.Equal(t, 100, *p1)

	// Case 2: Input is already *T
	valPtrT := &valT
	p2 := ptr.To[int](valPtrT)
	assert.NotNil(t, p2)
	assert.Equal(t, valPtrT, p2) // Should be the same pointer

	// Case 3: Input is a different type
	p3 := ptr.To[int]("a string")
	assert.NotNil(t, p3)
	assert.Equal(t, 0, *p3) // Should be a pointer to the zero value of int
}

func TestToVal(t *testing.T) {
	// Case 1: Input is a non-nil *T
	valT := 100
	valPtrT := &valT
	val1 := ptr.ToVal[int](valPtrT)
	assert.Equal(t, 100, val1)

	// Case 2: Input is a nil *T
	var nilPtr *int
	val2 := ptr.ToVal[int](nilPtr)
	assert.Equal(t, 0, val2)

	// Case 3: Input is of type T
	valT2 := 100
	val3 := ptr.ToVal[int](valT2)
	assert.Equal(t, 100, val3)

	// Case 4: Input is a different type
	val4 := ptr.ToVal[int]("a string")
	assert.Equal(t, 0, val4) // Should be the zero value of int
}
