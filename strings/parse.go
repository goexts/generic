// Package strings implements the functions, types, and interfaces for the module.
package strings

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// ParseOr converts string to specified type with default value support.
// Supported types: all basic types (int/uint variants, float, bool, string)
// Parameters:
//   - s: input string
//   - def: optional default value (returns first default value if conversion fails)
func ParseOr[T any](s string, def ...T) T {
	var t T
	switch p := any(&t).(type) {
	case *string:
		*p = s
		return t
	case *bool:
		v, err := strconv.ParseBool(s)
		if err == nil {
			*p = v
			return t
		}
	case *int:
		v, err := strconv.ParseInt(s, 10, 0)
		if err == nil {
			*p = int(v)
			return t
		}
	case *int8:
		v, err := strconv.ParseInt(s, 10, 8)
		if err == nil {
			*p = int8(v)
			return t
		}
	case *int16:
		v, err := strconv.ParseInt(s, 10, 16)
		if err == nil {
			*p = int16(v)
			return t
		}
	case *int32:
		v, err := strconv.ParseInt(s, 10, 32)
		if err == nil {
			*p = int32(v)
			return t
		}
	case *int64:
		v, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			*p = v
			return t
		}
	case *uint:
		v, err := strconv.ParseUint(s, 10, 0)
		if err == nil {
			*p = uint(v)
			return t
		}
	case *uint8:
		v, err := strconv.ParseUint(s, 10, 8)
		if err == nil {
			*p = uint8(v)
			return t
		}
	case *uint16:
		v, err := strconv.ParseUint(s, 10, 16)
		if err == nil {
			*p = uint16(v)
			return t
		}
	case *uint32:
		v, err := strconv.ParseUint(s, 10, 32)
		if err == nil {
			*p = uint32(v)
			return t
		}
	case *uint64:
		v, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			*p = v
			return t
		}
	case *float32:
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			*p = float32(v)
			return t
		}
	case *float64:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			*p = v
			return t
		}
	default:
		rt := reflect.TypeOf(t)
		if isJSONDeserializable(rt) {
			v, err := jsonUnmarshal[T](s)
			if err == nil {
				return v
			}
		}
	}

	if len(def) > 0 {
		return def[0]
	}

	panic(fmt.Sprintf("convert %q to %T failed", s, t))
}

func isJSONDeserializable(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		return true
	case reflect.Ptr:
		return isJSONDeserializable(rt.Elem())
	default:
		return false
	}
}

func jsonUnmarshal[T any](s string) (T, error) {
	var t T
	err := json.Unmarshal([]byte(s), &t)
	return t, err
}
