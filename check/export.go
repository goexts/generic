package check

// Includes returns the elements in src that are in ts.
func Includes[T comparable](src []T, ts ...T) []T {
	includes := make(map[T]struct{}, len(src))
	for _, s := range src {
		includes[s] = struct{}{}
	}
	var hits []T
	for _, t := range ts {
		if _, ok := includes[t]; ok {
			hits = append(hits, t)
		}
	}

	return hits
}

// Excludes returns the elements in src that are not in ts.
func Excludes[T comparable](src []T, ts ...T) []T {
	var excludes map[T]struct{}
	for _, s := range src {
		excludes[s] = struct{}{}
	}
	var hits []T
	for _, t := range ts {
		if _, ok := excludes[t]; !ok {
			hits = append(hits, t)
		}
	}
	return hits
}
