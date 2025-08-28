package cmp_test

import (
	"testing"

	"github.com/goexts/generic/cmp"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	assert.Equal(t, -1, cmp.Compare(1, 2))
	assert.Equal(t, 0, cmp.Compare(2, 2))
	assert.Equal(t, 1, cmp.Compare(2, 1))

	assert.Equal(t, -1, cmp.Compare("a", "b"))
	assert.Equal(t, 0, cmp.Compare("b", "b"))
	assert.Equal(t, 1, cmp.Compare("b", "a"))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, cmp.Min(1, 2))
	assert.Equal(t, 1, cmp.Min(2, 1))
	assert.Equal(t, 2, cmp.Min(2, 2))

	assert.Equal(t, 3.14, cmp.Min(3.14, 4.0))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, cmp.Max(1, 2))
	assert.Equal(t, 2, cmp.Max(2, 1))
	assert.Equal(t, 2, cmp.Max(2, 2))

	assert.Equal(t, "c", cmp.Max("a", "c"))
}

func TestClamp(t *testing.T) {
	// Value is within the range
	assert.Equal(t, 5, cmp.Clamp(5, 0, 10))

	// Value is less than the lower bound
	assert.Equal(t, 0, cmp.Clamp(-5, 0, 10))

	// Value is greater than the upper bound
	assert.Equal(t, 10, cmp.Clamp(15, 0, 10))

	// Value is equal to the bounds
	assert.Equal(t, 0, cmp.Clamp(0, 0, 10))
	assert.Equal(t, 10, cmp.Clamp(10, 0, 10))
}

func TestIsZero(t *testing.T) {
	assert.True(t, cmp.IsZero(0))
	assert.True(t, cmp.IsZero(""))
	assert.False(t, cmp.IsZero(1))
	assert.False(t, cmp.IsZero("hello"))

	var p *int
	assert.True(t, cmp.IsZero(p)) // nil pointer is the zero value
	p = new(int)
	assert.False(t, cmp.IsZero(p)) // non-nil pointer is not the zero value
}
