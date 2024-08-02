package types

type Array[T any] interface{ ~[...]T }

type Slice[T any] interface{ ~[]T }
