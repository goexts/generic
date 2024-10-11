package settings

// ApplyFunc is a ApplyFunc function for Apply
type ApplyFunc[S any] func(*S)

type ApplySetting[S any] interface {
	Apply(v *S)
}

type Setting[S any] interface {
	func(*S) | ApplyFunc[S]
}

func (s ApplyFunc[S]) Apply(v *S) {
	if v == nil {
		return
	}
	(s)(v)
}

// Apply is apply settings
func Apply[S any](d *S, ss []func(*S)) *S {
	if d == nil {
		return nil
	}
	for _, s := range ss {
		(s)(d)
	}
	return d
}

// ApplyOr is an apply settings with defaults
func ApplyOr[S any](s *S, ts ...func(*S)) *S {
	return Apply(s, ts)
}

// ApplyOrZero is an apply settings with defaults
func ApplyOrZero[S any](ss ...func(*S)) *S {
	var val S
	return Apply(&val, ss)
}
