package settings

// Setting is a setting function for Apply
type Setting[S any] func(*S)

// Apply is apply settings
func Apply[T Setting[S], S any](d S, ss ...T) *S {
	val := &d
	for _, s := range ss {
		s(val)
	}
	return val
}

// ApplyOrZero is an apply settings with defaults
func ApplyOrZero[T Setting[S], S any](ss ...T) *S {
	var val S
	for _, s := range ss {
		s(&val)
	}
	return &val
}
