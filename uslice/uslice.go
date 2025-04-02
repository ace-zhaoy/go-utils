package uslice

import (
	"sort"
)

func Intersect[T comparable](s1, s2 []T) []T {
	m, n := make(map[T]struct{}, len(s1)), make([]T, 0)
	for _, v := range s1 {
		m[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			n = append(n, v)
		}
	}
	return n
}

func Difference[T comparable](s1, s2 []T) []T {
	m := make(map[T]struct{}, len(s2))
	for _, item := range s2 {
		m[item] = struct{}{}
	}

	diff := make([]T, 0, len(s1))
	for _, item := range s1 {
		if _, found := m[item]; !found {
			diff = append(diff, item)
		}
	}

	return diff
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

// extractKey Go暂不支持接口合集，先隐式处理
func extractKey[K comparable, V any](v V, keyFunc func(V) K) K {
	if keyFunc != nil {
		return keyFunc(v)
	}
	if kv, ok := any(v).(interface{ GetID() K }); ok {
		return kv.GetID()
	}
	if kv, ok := any(v).(interface{ MapKey() K }); ok {
		return kv.MapKey()
	}
	if kv, ok := any(v).(interface{ GetID() any }); ok {
		if id, ok := kv.GetID().(K); ok {
			return id
		}
	}
	if kv, ok := any(v).(interface{ MapKey() any }); ok {
		if id, ok := kv.MapKey().(K); ok {
			return id
		}
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

func ToAnySlice[V any](s []V) []any {
	n := make([]any, len(s))
	for i, v := range s {
		n[i] = v
	}
	return n
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

func Range[V any](s []V, f func(item V) (finish bool)) {
	for _, v := range s {
		if f(v) {
			break
		}
	}
}

func Reverse[V any](s []V) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// SortSubsetByFullset sorts the elements of a subset slice in the order they appear in a fullset slice.
// Elements of the subset that do not exist in the fullset are moved to the end, maintaining their original order.
// This ensures a consistent order for subset elements based on the reference fullset,
// which is useful for synchronizing the order of items across different collections.
func SortSubsetByFullset[V comparable](fullset []V, subset []V) {
	positionMap := make(map[V]int, len(fullset))
	for i, val := range fullset {
		if _, ok := positionMap[val]; ok {
			continue
		}
		positionMap[val] = i
	}

	sort.SliceStable(subset, func(i, j int) bool {
		pi, ok := positionMap[subset[i]]
		if !ok {
			pi, positionMap[subset[i]] = len(fullset), len(fullset)
		}
		pj, ok := positionMap[subset[j]]
		if !ok {
			pj, positionMap[subset[j]] = len(fullset), len(fullset)
		}
		return pi < pj
	})
}

// AlignOtherSetsByNewSubsetOrder adjusts the order of elements in other sets based on a new subset order.
// This function takes three parameters:
// - oldSubsetIndexMap, which maps elements to their original indices in the subset.
// - newSubset, which contains the elements in their new desired order.
// - otherSets, which is a variadic parameter of slices that should be reordered based on the newSubset's order.
func AlignOtherSetsByNewSubsetOrder[V comparable, S any](oldSubsetIndexMap map[V][]int, newSubset []V, otherSets ...[]S) {
	subsetIndexSwapMap := make(map[int]int, len(newSubset))
	for i, v := range newSubset {
		oldIndex := oldSubsetIndexMap[v][0]
		oldSubsetIndexMap[v] = oldSubsetIndexMap[v][1:]
		if oldIndex == i {
			continue
		}
		subsetIndexSwapMap[oldIndex] = i
	}

	jumpSet := make(map[int]struct{}, len(subsetIndexSwapMap))
	for i := range otherSets {
		if len(otherSets[i]) != len(newSubset) {
			continue
		}
		for j := range otherSets[i] {
			if _, ok := jumpSet[j]; ok {
				continue
			}
			ni, ok := subsetIndexSwapMap[j]
			if !ok {
				continue
			}
			otherSets[i][j], otherSets[i][ni] = otherSets[i][ni], otherSets[i][j]
			for {
				bi, ok := subsetIndexSwapMap[ni]
				if !ok {
					break
				}
				jumpSet[ni] = struct{}{}
				if j == bi {
					break
				}
				otherSets[i][j], otherSets[i][bi], ni = otherSets[i][bi], otherSets[i][j], bi
			}
		}
	}
}

// ElementIndicesMap creates a mapping from each unique element in the slice to a list of its indices.
func ElementIndicesMap[V comparable](s []V) map[V][]int {
	m := make(map[V][]int, len(s))
	for i := range s {
		m[s[i]] = append(m[s[i]], i)
	}
	return m
}

// SortSubsetsByFullset sorts a subset according to the order in the fullset and adjusts the positions
// of corresponding elements in other sets accordingly. Duplicate elements are ordered by their first occurrence position.
func SortSubsetsByFullset[V comparable, S any](fullset []V, subset []V, otherSets ...[]S) {
	oldSubsetIndexMap := ElementIndicesMap(subset)

	SortSubsetByFullset(fullset, subset)

	AlignOtherSetsByNewSubsetOrder(oldSubsetIndexMap, subset, otherSets...)
}

// SortSubsetByFullsetOrder sorts the subset completely according to the order in the fullset.
// This function takes two slices: `fullset` and `subset`. Elements in `subset` are rearranged to match the order they appear in `fullset`.
// If there are elements in `subset` that do not exist in `fullset`, they are moved to the end of `subset` in their original order.
func SortSubsetByFullsetOrder[V comparable](fullset []V, subset []V) {
	subsetIndexMap := ElementIndicesMap(subset)

	i := 0
	for j := range fullset {
		if _, ok := subsetIndexMap[fullset[j]]; !ok {
			continue
		}
		subset[i] = fullset[j]
		i++
		if len(subsetIndexMap[fullset[j]]) == 1 {
			delete(subsetIndexMap, fullset[j])
		} else {
			subsetIndexMap[fullset[j]] = subsetIndexMap[fullset[j]][1:]
		}
	}
	if len(subsetIndexMap) > 0 {
		unknownMap := make(map[int]V, len(subsetIndexMap))
		for v, indices := range subsetIndexMap {
			for _, index := range indices {
				unknownMap[index] = v
			}
		}
		for j := 0; j < len(subset); j++ {
			if v, ok := unknownMap[j]; ok {
				subset[i] = v
				i++
			}
		}
	}
}

// SortSubsetsByFullsetOrder rearranges the elements of a subset to match the order of a fullset
// and adjusts the order of elements in other associated sets based on the new order of the subset.
// This function takes the following parameters:
// - fullset, which is the reference slice containing elements in the desired order.
// - subset, which is the slice that needs to be reordered to match the order of elements in fullset.
// - otherSets, which is a variadic parameter of slices that should be reordered according to the new order of the subset.
func SortSubsetsByFullsetOrder[V comparable, S any](fullset []V, subset []V, otherSets ...[]S) {
	oldSubsetIndexMap := ElementIndicesMap(subset)

	SortSubsetByFullsetOrder(fullset, subset)
	AlignOtherSetsByNewSubsetOrder(oldSubsetIndexMap, subset, otherSets...)
}
