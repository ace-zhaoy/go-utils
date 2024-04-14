package ureflect

import "reflect"

func HasField(s any, name string) bool {
	v := reflect.TypeOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	_, e := v.FieldByName(name)
	return e
}

func FieldExists[T any](name string) bool {
	var t T
	return HasField(t, name)
}
