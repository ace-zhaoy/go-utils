package utype

func Pointer[T any](p T) *T {
	return &p
}

func Dereference[T any](p *T) T {
	if p != nil {
		return *p
	}
	var t T
	return t
}

func DereferenceOrDefault[T any](p *T, defaultValue T) T {
	if p != nil {
		return *p
	}
	return defaultValue
}
