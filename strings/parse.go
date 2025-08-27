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
	rt := reflect.TypeOf(t)
	switch rt.Kind() {
	case reflect.String:
		return any(s).(T)
	case reflect.Bool:
		v, err := strconv.ParseBool(s)
		if err == nil {
			return any(v).(T)
		}
	case reflect.Int:
		v, err := strconv.ParseInt(s, 10, 0)
		if err == nil {
			return any(int(v)).(T)
		}
	case reflect.Int8:
		v, err := strconv.ParseInt(s, 10, 8)
		if err == nil {
			return any(int8(v)).(T)
		}
	case reflect.Int16:
		v, err := strconv.ParseInt(s, 10, 16)
		if err == nil {
			return any(int16(v)).(T)
		}
	case reflect.Int32:
		v, err := strconv.ParseInt(s, 10, 32)
		if err == nil {
			return any(int32(v)).(T)
		}
	case reflect.Int64:
		v, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return any(v).(T)
		}
	case reflect.Uint:
		v, err := strconv.ParseUint(s, 10, 0)
		if err == nil {
			return any(uint(v)).(T)
		}
	case reflect.Uint8:
		v, err := strconv.ParseUint(s, 10, 8)
		if err == nil {
			return any(uint8(v)).(T)
		}
	case reflect.Uint16:
		v, err := strconv.ParseUint(s, 10, 16)
		if err == nil {
			return any(uint16(v)).(T)
		}
	case reflect.Uint32:
		v, err := strconv.ParseUint(s, 10, 32)
		if err == nil {
			return any(uint32(v)).(T)
		}
	case reflect.Uint64:
		v, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			return any(v).(T)
		}
	case reflect.Float32:
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return any(float32(v)).(T)
		}
	case reflect.Float64:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return any(v).(T)
		}
	default:
		if isJSONDeserializable(rt) {
			v, err := jsonUnmarshal[T](s, rt)
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
	}
	return false
}

func jsonUnmarshal[T any](s string, rt reflect.Type) (T, error) {
	var t T
	target := any(t)
	if rt.Kind() != reflect.Ptr {
		target = &t
	}
	err := json.Unmarshal([]byte(s), target)
	return t, err
}
