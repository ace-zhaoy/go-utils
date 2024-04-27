package umap

func Foreach[K comparable, V any](m map[K]V, f func(k K, v V)) {
	for k, v := range m {
		f(k, v)
	}
}

func KeyExists[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func FindKeyByValue[K, V comparable](m map[K]V, v V) (k K, ok bool) {
	for k1, v1 := range m {
		if v1 == v {
			return k1, true
		}
	}
	return
}
