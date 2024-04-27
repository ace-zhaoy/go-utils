package ureflect

import (
	"reflect"
)

func HasField[T any](value T, fieldName string) bool {
	val := reflect.ValueOf(value)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			typeOfT := reflect.TypeOf(value).Elem()
			if typeOfT.Kind() == reflect.Struct {
				_, ok := typeOfT.FieldByName(fieldName)
				return ok
			}
		}
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		fieldVal := val.FieldByName(fieldName)
		return fieldVal.IsValid()
	}

	return false
}

type Zeroable interface {
	IsZero() bool
}

func IsZero[T any](val T) bool {
	if zeroable, ok := any(val).(Zeroable); ok {
		return zeroable.IsZero()
	}
	v := reflect.ValueOf(val)
	if (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && v.IsNil() {
		return true
	}
	if v.Kind() == reflect.Invalid {
		return true
	}

	return v.IsZero()
}

func IsRealZero[T any](val T) bool {
	if zeroable, ok := any(val).(Zeroable); ok {
		return zeroable.IsZero()
	}

	v := reflect.ValueOf(val)
	for {
		switch v.Kind() {
		case reflect.Ptr, reflect.Interface:
			if v.IsNil() {
				return true
			}
			v = v.Elem()
			continue
		case reflect.Invalid:
			return true
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
			return v.Len() == 0
		default:
			return v.IsZero()
		}
	}
}
