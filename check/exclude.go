package check

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
