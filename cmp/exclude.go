package cmp

// Excludes returns the elements in src that are not in ts.
func Excludes[T comparable](src []T, ts ...T) []T {
	excludes := make(map[T]struct{}, len(src))
	for _, t := range src {
		excludes[t] = struct{}{}
	}
	var hits []T
	for _, t := range ts {
		if _, ok := excludes[t]; !ok {
			hits = append(hits, t)
		}
	}
	return hits
}
