package uslice

func Intersect[T comparable](s1, s2 []T) []T {
	m, n := make(map[T]struct{}, len(s1)), make([]T, 0)
	for _, v := range s1 {
		m[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			n = append(n, v)
			delete(m, v)
		}
	}
	return n
}

func Difference[T comparable](s1, s2 []T) []T {
	m, n := make(map[T]struct{}, len(s1)), make([]T, 0)
	for _, v := range s1 {
		m[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := m[v]; !ok {
			n = append(n, v)
		}
	}
	return n
}

func Union[T comparable](ss ...[]T) []T {
	m, n := make(map[T]struct{}), make([]T, 0)
	for _, s := range ss {
		for _, v := range s {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				n = append(n, v)
			}
		}
	}
	return n
}

func Contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func Unique[T comparable](s []T) []T {
	m, n := make(map[T]struct{}, len(s)), make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			n = append(n, v)
		}
	}
	return n
}

func ToSet[T comparable](s []T) map[T]struct{} {
	m := make(map[T]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

func Chunk[T any](s []T, length uint) [][]T {
	n := make([][]T, 0)
	sl := len(s)
	if length == 0 || sl == 0 {
		return n
	}

	for i := 0; i < len(s); i += int(length) {
		end := i + int(length)
		if end > len(s) {
			end = len(s)
		}
		n = append(n, s[i:end])
	}

	return n
}

func extractKey[K comparable, V any](v V, keyFunc func(V) K) K {
	if keyFunc != nil {
		return keyFunc(v)
	}
	if kv, ok := any(v).(interface{ GetID() any }); ok {
		return kv.GetID().(K)
	}
	if kv, ok := any(v).(interface{ GetID() K }); ok {
		return kv.GetID()
	}
	if kv, ok := any(v).(interface{ MapKey() any }); ok {
		return kv.MapKey().(K)
	}
	if kv, ok := any(v).(interface{ MapKey() K }); ok {
		return kv.MapKey()
	}
	panic("extractKey: no key function provided and value doesn't implement GetID or MapKey")
}

func ToMap[K comparable, V any](s []V) map[K]V {
	m := make(map[K]V, len(s))
	for _, v := range s {
		m[extractKey[K](v, nil)] = v
	}
	return m
}

func ToMapF[K comparable, V any](s []V, keyFunc func(V) K) map[K]V {
	m := make(map[K]V, len(s))
	for _, v := range s {
		m[extractKey(v, keyFunc)] = v
	}
	return m
}

func ToMapV[K comparable, V, T any](s []T, valueFunc func(T) V) map[K]V {
	m := make(map[K]V, len(s))
	for _, v := range s {
		m[extractKey[K](v, nil)] = valueFunc(v)
	}
	return m
}

func ToMapFV[K comparable, V, T any](s []T, keyFunc func(T) K, valueFunc func(T) V) map[K]V {
	m := make(map[K]V, len(s))
	for _, v := range s {
		m[extractKey(v, keyFunc)] = valueFunc(v)
	}
	return m
}

func ToSliceMap[K comparable, V any](s []V, keyFunc func(V) K) map[K][]V {
	m := make(map[K][]V, len(s))
	for _, v := range s {
		k := extractKey(v, keyFunc)
		m[k] = append(m[k], v)
	}
	return m
}

func ToSliceMapV[K comparable, V, T any](s []T, keyFunc func(T) K, valueFunc func(T) V) map[K][]V {
	m := make(map[K][]V, len(s))
	for _, v := range s {
		k := extractKey(v, keyFunc)
		m[k] = append(m[k], valueFunc(v))
	}
	return m
}

func Copy[V any](seq []V, maxLen int) []V {
	if maxLen <= 0 {
		return make([]V, 0)
	}
	if len(seq) <= maxLen {
		return append([]V(nil), seq...)
	}
	return append([]V(nil), seq[:maxLen]...)
}

func MapI[V, T any](s []V, f func(index int, item V) T) []T {
	n := make([]T, len(s))
	for i, v := range s {
		n[i] = f(i, v)
	}
	return n
}

func Map[V, T any](s []V, f func(item V) T) []T {
	return MapI(s, func(_ int, item V) T { return f(item) })
}

func FilterI[V any](s []V, f func(index int, item V) bool) []V {
	n := make([]V, 0, len(s))
	for i, v := range s {
		if f(i, v) {
			n = append(n, v)
		}
	}
	return n
}

func Filter[V any](s []V, f func(item V) bool) []V {
	return FilterI(s, func(_ int, item V) bool { return f(item) })
}

func ReduceI[V, T any](s []V, f func(index int, item V, accumulator T) T, initial ...T) T {
	var accumulator T
	if len(initial) > 0 {
		accumulator = initial[0]
	}

	for i, v := range s {
		accumulator = f(i, v, accumulator)
	}
	return accumulator
}

func Reduce[V, T any](s []V, f func(item V, accumulator T) T, initial ...T) T {
	return ReduceI(s, func(_ int, item V, accumulator T) T { return f(item, accumulator) }, initial...)
}

func ForEachI[V any](s []V, f func(index int, item V)) {
	for i, v := range s {
		f(i, v)
	}
}

func ForEach[V any](s []V, f func(item V)) {
	ForEachI(s, func(_ int, item V) { f(item) })
}

func Reverse[V any](s []V) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
