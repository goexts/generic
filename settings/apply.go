package settings

// Setting is a setting function for Apply
type Setting[S any] func(*S)

func (s *Setting[S]) Apply(v *S) {
	if v == nil {
		return
	}
	(*s)(v)
}

// Apply is apply settings
func Apply[T Setting[S], S any](d *S, ss []T) *S {
	if d == nil {
		return nil
	}
	for _, s := range ss {
		s(d)
	}
	return d
}

// ApplyOr is an apply settings with defaults
func ApplyOr[T Setting[S], S any](d S, ss ...T) *S {
	return Apply(&d, ss)
}

// ApplyOrZero is an apply settings with defaults
func ApplyOrZero[T Setting[S], S any](ss ...T) *S {
	var val S
	return Apply(&val, ss)
}
