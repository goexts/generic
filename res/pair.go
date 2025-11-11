package res

// Pair represents an immutable pair of two different types A and B.
// It's useful for combining two values of different types into a single value
// that can be passed around and manipulated as a unit.
type Pair[A, B any] struct {
	first  A
	second B
}

// NewPair creates a new Pair with the given values.
func NewPair[A, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{
		first:  first,
		second: second,
	}
}

// First returns the first value of the pair.
func (p Pair[A, B]) First() A {
	return p.first
}

// Second returns the second value of the pair.
func (p Pair[A, B]) Second() B {
	return p.second
}

// Values returns both values from the pair.
func (p Pair[A, B]) Values() (A, B) {
	return p.first, p.second
}

// Map applies the given functions to the pair's values and returns a new pair.
// This allows transforming both values in the pair in a single operation.
func (p Pair[A, B]) Map(fnA func(A) A, fnB func(B) B) Pair[A, B] {
	return Pair[A, B]{
		first:  fnA(p.first),
		second: fnB(p.second),
	}
}

// MapFirst applies the given function to the first value of the pair.
// The second value remains unchanged.
func (p Pair[A, B]) MapFirst(fn func(A) A) Pair[A, B] {
	return Pair[A, B]{
		first:  fn(p.first),
		second: p.second,
	}
}

// MapSecond applies the given function to the second value of the pair.
// The first value remains unchanged.
func (p Pair[A, B]) MapSecond(fn func(B) B) Pair[A, B] {
	return Pair[A, B]{
		first:  p.first,
		second: fn(p.second),
	}
}

// Swap returns a new Pair with the values swapped.
func (p Pair[A, B]) Swap() Pair[B, A] {
	return Pair[B, A]{
		first:  p.second,
		second: p.first,
	}
}

// WithFirst returns a new Pair with the first value replaced by the given value.
func (p Pair[A, B]) WithFirst(first A) Pair[A, B] {
	return Pair[A, B]{
		first:  first,
		second: p.second,
	}
}

// WithSecond returns a new Pair with the second value replaced by the given value.
func (p Pair[A, B]) WithSecond(second B) Pair[A, B] {
	return Pair[A, B]{
		first:  p.first,
		second: second,
	}
}
