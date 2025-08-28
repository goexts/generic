package set

// Contains checks if a slice `s` contains the element `e`.
// The check is performed using the equality operator (==).
func Contains[T comparable](s []T, e T) bool {
	for _, t := range s {
		if t == e {
			return true
		}
	}
	return false
}

// Exists checks if at least one element in a slice `s` satisfies the predicate `f`.
func Exists[T any](s []T, f func(T) bool) bool {
	for _, t := range s {
		if f(t) {
			return true
		}
	}
	return false
}

// Unique returns a new slice containing only the unique elements of the input
// slice `s`. The order of elements in the returned slice is not guaranteed.
func Unique[T comparable](s []T) []T {
	set := make(map[T]struct{}, len(s))
	for _, t := range s {
		set[t] = struct{}{}
	}

	result := make([]T, 0, len(set))
	for t := range set {
		result = append(result, t)
	}
	return result
}

// Union returns a new slice containing the unique elements present in either
// s1 or s2. The order of elements in the returned slice is not guaranteed.
func Union[T comparable](s1, s2 []T) []T {
	set := make(map[T]struct{}, len(s1)+len(s2))
	for _, t := range s1 {
		set[t] = struct{}{}
	}
	for _, t := range s2 {
		set[t] = struct{}{}
	}

	result := make([]T, 0, len(set))
	for t := range set {
		result = append(result, t)
	}
	return result
}

// Intersection returns a new slice containing only the elements that exist in
// both s1 and s2. The order of elements in the returned slice is not guaranteed.
func Intersection[T comparable](s1, s2 []T) []T {
	set := make(map[T]struct{}, len(s1))
	for _, t := range s1 {
		set[t] = struct{}{}
	}

	var result []T
	for _, t := range s2 {
		if _, ok := set[t]; ok {
			result = append(result, t)
		}
	}

	return result
}

// Difference returns a new slice containing the elements from s2 that are not
// present in s1 (s2 - s1). The order of elements in the returned slice is not guaranteed.
func Difference[T comparable](s1, s2 []T) []T {
	excludeSet := make(map[T]struct{}, len(s1))
	for _, item := range s1 {
		excludeSet[item] = struct{}{}
	}

	resultSet := make(map[T]struct{})
	for _, item := range s2 {
		if _, found := excludeSet[item]; !found {
			resultSet[item] = struct{}{}
		}
	}

	result := make([]T, 0, len(resultSet))
	for item := range resultSet {
		result = append(result, item)
	}
	return result
}
