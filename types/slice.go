package types

// remove because it is not supported
// type Array[T any] interface{ ~[...]T }

type Slice[T any] interface{ ~[]T }
