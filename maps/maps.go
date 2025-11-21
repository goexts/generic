package maps

// FirstKey searches for the first key from the provided list that exists in the given map.
// It returns the found key and true if a key is found, otherwise it returns the zero value of K and false.
//
// Parameters:
//
//	m: The map to search within.
//	keys: A variadic list of keys to look for in the map.
//
// Returns:
//
//	The first key from the 'keys' list that exists in 'm', and true.
//	If no key from 'keys' exists in 'm', it returns the zero value of K and false.
func FirstKey[K comparable, V any](m map[K]V, keys ...K) (K, bool) {
	for _, key := range keys {
		if _, ok := m[key]; ok {
			return key, true
		}
	}
	var zeroK K // Return zero value of K if no key is found
	return zeroK, false
}

// FirstKeyBy searches for the first key from the provided list that,
// after being transformed by the 'transform' function, exists in the given map.
// It returns the original key from the list and true if a key is found,
// otherwise it returns the zero value of ListK and false.
// This is useful when the keys in the list need a special comparison or normalization
// before being looked up in the map (e.g., case-insensitive string comparison).
//
// Parameters:
//
//	m: The map to search within. The keys of this map are of type MapK.
//	transform: A function that converts a key of type ListK to a key of type MapK.
//	keys: A variadic list of keys (of type ListK) to look for in the map after transformation.
//
// Returns:
//
//	The first key from the 'keys' list (of type ListK) whose transformed value exists in 'm', and true.
//	If no transformed key from 'keys' exists in 'm', it returns the zero value of ListK and false.
func FirstKeyBy[ListK any, MapK comparable, V any](m map[MapK]V, transform func(ListK) MapK, keys ...ListK) (ListK, bool) {
	for _, key := range keys {
		transformedKey := transform(key)
		if _, ok := m[transformedKey]; ok {
			return key, true
		}
	}
	var zeroListK ListK // Return zero value of ListK if no key is found
	return zeroListK, false
}
