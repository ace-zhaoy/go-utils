package ucompare

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](args ...T) T {
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

func Min[T constraints.Ordered](args ...T) T {
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
