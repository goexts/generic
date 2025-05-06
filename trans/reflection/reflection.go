/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package reflection implements the functions, types, and interfaces for the module.
package reflection

import (
	"fmt"
	"reflect"
)

type InvalidStructError struct {
	Type reflect.Type
}

func NewInvalidStructError(v interface{}) error {
	return &InvalidStructError{Type: reflect.TypeOf(v)}
}

func (e *InvalidStructError) Error() string {
	return fmt.Sprintf("expected struct type, got %s", e.Type)
}

type TypeMismatchError struct {
	Actual, Expected reflect.Type
}

func NewTypeMismatchError(actual, expected reflect.Type) error {
	return &TypeMismatchError{Actual: actual, Expected: expected}
}

func (e *TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch: expected %s, got %s", e.Expected, e.Actual)
}

type FieldNotFoundError struct {
	TargetType, StructType reflect.Type
}

func NewFieldNotFoundError(target, strct reflect.Type) error {
	return &FieldNotFoundError{TargetType: target, StructType: strct}
}

func (e *FieldNotFoundError) Error() string {
	return fmt.Sprintf("field of type %s not found in %s", e.TargetType, e.StructType)
}

func fieldValueByType[T any](v reflect.Value, targetType reflect.Type) (interface{}, error) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := field.Type()

		if fieldType == targetType {
			return field.Interface(), nil
		}

		if field.Kind() == reflect.Struct {
			if result, err := fieldValueByType[T](field, targetType); err == nil {
				return result, nil
			}
		}
	}

	return nil, NewFieldNotFoundError(targetType, v.Type())
}

func FieldValueByType[T any](obj T) (T, error) {
	var zero T
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return zero, NewInvalidStructError(obj)
	}

	value, err := fieldValueByType[T](val, reflect.TypeOf(zero))
	if err != nil {
		return zero, err
	}

	if result, ok := value.(T); ok {
		return result, nil
	}

	return zero, NewTypeMismatchError(val.Type(), reflect.TypeOf(zero))
}

// FieldPointByType returns the first field pointer in the struct that matches the target type
func FieldPointByType[T any](obj T) (*T, error) {
	value, err := FieldValueByType(obj)
	if err != nil {
		return nil, err
	}
	return &value, nil
}
