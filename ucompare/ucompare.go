package ucompare

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Max[T Ordered](args ...T) T {
	l := len(args)
	var t T
	if l > 0 {
		t = args[0]
	}

	for i := 1; i < l; i++ {
		if args[i] > t {
			t = args[i]
		}
	}

	return t
}

func Min[T Ordered](args ...T) T {
	l := len(args)
	var t T
	if l > 0 {
		t = args[0]
	}

	for i := 1; i < l; i++ {
		if args[i] < t {
			t = args[i]
		}
	}

	return t
}
