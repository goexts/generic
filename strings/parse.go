// Package strings implements the functions, types, and interfaces for the module.
package strings

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// ParseOr converts a string to a specified type, with support for a default value.
// It supports all basic types (int/uint variants, float, bool, string) and JSON-deserializable types.
// If parsing fails and a default value is provided, it returns the default value.
// If parsing fails and no default is provided, it panics.
func ParseOr[T any](s string, def ...T) T {
	var t T
	rt := reflect.TypeOf(t)

	var val any
	var err error

	switch rt.Kind() {
	case reflect.String:
		// For string types, no conversion is needed.
		// We must cast `s` to `any` first before casting to the generic type `T`.
		return any(s).(T) //nolint:errcheck
	case reflect.Bool:
		val, err = strconv.ParseBool(s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// Use the bit size of the target type for correct parsing.
		bitSize := rt.Bits()
		val, err = strconv.ParseInt(s, 10, bitSize)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		bitSize := rt.Bits()
		val, err = strconv.ParseUint(s, 10, bitSize)
	case reflect.Float32, reflect.Float64:
		bitSize := rt.Bits()
		val, err = strconv.ParseFloat(s, bitSize)
	default:
		// For complex types, attempt to unmarshal from JSON.
		if isJSONDeserializable(rt) {
			v, jsonErr := jsonUnmarshal[T](s)
			if jsonErr == nil {
				return v
			}
		}
		// If JSON fails or the type is not supported, fall through to the default/panic logic.
		err = fmt.Errorf("unsupported type for parsing: %T", t)
	}

	// If any parsing error occurred...
	if err != nil {
		if len(def) > 0 {
			return def[0]
		}
		panic(fmt.Sprintf("convert %q to %T failed: %v", s, t, err))
	}

	// If parsing was successful, convert the parsed value (which is int64, uint64, float64, or bool)
	// to the actual target type T.
	return reflect.ValueOf(val).Convert(rt).Interface().(T) //nolint:errcheck
}

// isJSONDeserializable checks if a type is likely to be unmarshaled from JSON.
func isJSONDeserializable(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		return true
	case reflect.Ptr:
		// Recurse on pointer element types.
		return isJSONDeserializable(rt.Elem())
	default:
		return false
	}
}

// jsonUnmarshal is a helper to unmarshal a string into a generic type.
func jsonUnmarshal[T any](s string) (T, error) {
	var t T
	err := json.Unmarshal([]byte(s), &t)
	return t, err
}
