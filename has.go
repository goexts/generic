package generic

// Has checks whether a value is in a slice.
func Has[T comparable](s []T, e T) bool {
	var t T
	for _, t = range s {
		if t == e {
			return true
		}
	}
	return false
}

// HasFunc checks whether a value is in a slice using a function.
func HasFunc[T any](s []T, f func(T) bool) bool {
	for _, t := range s {
		if f(t) {
			return true
		}
	}
	return false
}

// NotIn checks whether a value is not in a slice.
func NotIn[T comparable](s []T, e T) bool {
	return !Has(s, e)
}
