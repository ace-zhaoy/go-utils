package ucondition

import "github.com/ace-zhaoy/go-utils/ureflect"

func If[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func IfF[T any](condition bool, trueFunc func() T, falseFunc func() T) T {
	if condition {
		return trueFunc()
	}
	return falseFunc()
}

func IfTF[T any](condition bool, trueFunc func() T, falseValue T) T {
	if condition {
		return trueFunc()
	}
	return falseValue
}

func IfFF[T any](condition bool, trueValue T, falseFunc func() T) T {
	if condition {
		return trueValue
	}
	return falseFunc()
}

func IfZero[T any](value T, defaultValue T) T {
	if ureflect.IsZero(value) {
		return defaultValue
	}
	return value
}

func IfZeroF[T any](value T, defaultValueFunc func() T) T {
	if ureflect.IsZero(value) {
		return defaultValueFunc()
	}
	return value
}
