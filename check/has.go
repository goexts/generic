package check

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
