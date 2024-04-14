package utype

func Pointer[T any](p T) *T {
	return &p
}
