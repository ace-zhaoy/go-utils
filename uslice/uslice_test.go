package uslice

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args[T comparable] struct {
		s1 []T
		s2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s1: []int{1, 2, 3, 4, 5},
				s2: []int{2, 4, 6},
			},
			want: []int{2, 4},
		},
		{
			name: "test2",
			args: args[int]{
				s1: []int{1, 2, 3},
				s2: []int{4, 5, 6},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type args[T comparable] struct {
		s1 []T
		s2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s1: []int{1, 2, 3, 4, 5},
				s2: []int{2, 4, 6},
			},
			want: []int{6},
		},
		{
			name: "test2",
			args: args[int]{
				s1: []int{1, 2, 3},
				s2: []int{4, 5, 6},
			},
			want: []int{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args[T comparable] struct {
		ss [][]T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				ss: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test2",
			args: args[int]{
				ss: [][]int{{1, 2}, {3, 4}, {5, 6}},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.ss...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		s []T
		v T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
				v: 3,
			},
			want: true,
		},
		{
			name: "test2",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
				v: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args[T comparable] struct {
		s []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5, 3, 4},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test2",
			args: args[int]{
				s: []int{},
			},
			want: []int{},
		},
		{
			name: "test3",
			args: args[int]{
				s: []int{1, 1, 1, 2, 2, 3, 3, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSet(t *testing.T) {
	type args[T comparable] struct {
		s []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want map[T]struct{}
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s: []int{1, 2, 3, 2, 1},
			},
			want: map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSet(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args[T any] struct {
		s      []T
		length uint
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s:      []int{1, 2, 3, 4, 5},
				length: 2,
			},
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "test2",
			args: args[int]{
				s:      []int{1, 2, 3, 4, 5, 6},
				length: 3,
			},
			want: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.s, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

type M1 struct {
	ID int
}

func (m *M1) GetID() any {
	return m.ID
}

func TestToMap(t *testing.T) {
	type args[T comparable] struct {
		s []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want map[int]T
	}

	tests := []testCase[*M1]{
		{
			name: "test1",
			args: args[*M1]{
				s: []*M1{
					{ID: 1},
					{ID: 2},
					{ID: 3},
				},
			},
			want: map[int]*M1{
				1: {ID: 1},
				2: {ID: 2},
				3: {ID: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMap[int](tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapF(t *testing.T) {
	type args[V any, K comparable] struct {
		s       []V
		keyFunc func(V) K
	}
	type testCase[V any, K comparable] struct {
		name string
		args args[V, K]
		want map[K]V
	}
	tests := []testCase[int, string]{
		{
			name: "test1",
			args: args[int, string]{
				s:       []int{1, 2, 3},
				keyFunc: func(item int) string { return strconv.Itoa(item) },
			},
			want: map[string]int{"1": 1, "2": 2, "3": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapF(tt.args.s, tt.args.keyFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapV(t *testing.T) {
	type args[T any, K comparable, V any] struct {
		s         []T
		keyFunc   func(T) K
		valueFunc func(T) V
	}
	type testCase[T any, K comparable, V any] struct {
		name string
		args args[T, K, V]
		want map[K]V
	}
	tests := []testCase[int, string, string]{
		{
			name: "test1",
			args: args[int, string, string]{
				s:         []int{1, 2, 3},
				keyFunc:   func(item int) string { return strconv.Itoa(item) },
				valueFunc: func(item int) string { return fmt.Sprintf("value %d", item) },
			},
			want: map[string]string{"1": "value 1", "2": "value 2", "3": "value 3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapFV(tt.args.s, tt.args.keyFunc, tt.args.valueFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSliceMap(t *testing.T) {
	type args[K comparable, V any] struct {
		s       []V
		keyFunc func(V) K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K][]V
	}
	tests := []testCase[int, string]{
		{
			name: "test1",
			args: args[int, string]{
				s: []string{"apple", "banana", "cherry", "date", "apple"},
				keyFunc: func(s string) int {
					return len(s)
				},
			},
			want: map[int][]string{
				5: {"apple", "apple"},
				6: {"banana", "cherry"},
				4: {"date"},
			},
		},
		{
			name: "test2",
			args: args[int, string]{
				s:       []string{},
				keyFunc: func(s string) int { return len(s) },
			},
			want: map[int][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSliceMap(tt.args.s, tt.args.keyFunc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSliceMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSliceMapV(t *testing.T) {
	type args[K comparable, V, T any] struct {
		s         []T
		keyFunc   func(T) K
		valueFunc func(T) V
	}
	type testCase[K comparable, V, T any] struct {
		name string
		args args[K, V, T]
		want map[K][]V
	}
	tests := []testCase[int, string, int]{
		{
			name: "test1",
			args: args[int, string, int]{
				s: []int{1, 2, 3, 4, 5},
				keyFunc: func(i int) int {
					return i % 2
				},
				valueFunc: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: map[int][]string{
				1: {"1", "3", "5"},
				0: {"2", "4"},
			},
		},
		{
			name: "test2",
			args: args[int, string, int]{
				s:         []int{},
				keyFunc:   func(i int) int { return i % 2 },
				valueFunc: func(i int) string { return strconv.Itoa(i) },
			},
			want: map[int][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSliceMapV(tt.args.s, tt.args.keyFunc, tt.args.valueFunc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSliceMapV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args[V any] struct {
		seq    []V
		maxLen int
	}
	type testCase[V any] struct {
		name string
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				seq:    []int{1, 2, 3, 4, 5},
				maxLen: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test2",
			args: args[int]{
				seq:    []int{1, 2, 3, 4, 5},
				maxLen: 10,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy(tt.args.seq, tt.args.maxLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapI(t *testing.T) {
	type args[V any, T any] struct {
		s []V
		f func(index int, item V) T
	}
	type testCase[V any, T any] struct {
		name string
		args args[V, T]
		want []T
	}
	tests := []testCase[int, string]{
		{
			name: "test1",
			args: args[int, string]{
				s: []int{1, 2, 3},
				f: func(index int, item int) string { return fmt.Sprintf("%d: %d", index, item) },
			},
			want: []string{"0: 1", "1: 2", "2: 3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapI(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[V any, T any] struct {
		s []V
		f func(item V) T
	}
	type testCase[V any, T any] struct {
		name string
		args args[V, T]
		want []T
	}
	tests := []testCase[int, string]{
		{
			name: "test1",
			args: args[int, string]{
				s: []int{1, 2, 3},
				f: func(item int) string { return strconv.Itoa(item) },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterI(t *testing.T) {
	type args[V any] struct {
		s []V
		f func(index int, item V) bool
	}
	type testCase[V any] struct {
		name string
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
				f: func(index int, item int) bool { return index%2 == 0 },
			},
			want: []int{1, 3, 5},
		},
		{
			name: "test2",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
				f: func(index int, item int) bool { return item > 3 },
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterI(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[V any] struct {
		s []V
		f func(item V) bool
	}
	type testCase[V any] struct {
		name string
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
				f: func(item int) bool { return item%2 == 0 },
			},
			want: []int{2, 4},
		},
		{
			name: "test2",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
				f: func(item int) bool { return item > 3 },
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceI(t *testing.T) {
	type args[V any, T any] struct {
		s       []V
		f       func(index int, item V, accumulator T) T
		initial []T
	}
	type testCase[V any, T any] struct {
		name string
		args args[V, T]
		want T
	}
	tests := []testCase[int, int]{
		{
			name: "test1",
			args: args[int, int]{
				s:       []int{1, 2, 3, 4, 5},
				f:       func(index int, item int, accumulator int) int { return accumulator + item + index },
				initial: []int{0},
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceI(tt.args.s, tt.args.f, tt.args.initial...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReduceI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args[V any, T any] struct {
		s       []V
		f       func(item V, accumulator T) T
		initial []T
	}
	type testCase[V any, T any] struct {
		name string
		args args[V, T]
		want T
	}
	tests := []testCase[int, int]{
		{
			name: "test1",
			args: args[int, int]{
				s:       []int{1, 2, 3, 4, 5},
				f:       func(item int, accumulator int) int { return accumulator + item },
				initial: []int{0},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.s, tt.args.f, tt.args.initial...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	type args[V any] struct {
		s []V
		f func(item V)
	}
	type testCase[V any] struct {
		name string
		args args[V]
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
				f: func(item int) { t.Logf("%d", item) },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForEach(tt.args.s, tt.args.f)
		})
	}
}

func TestForEachI(t *testing.T) {
	type args[V any] struct {
		s []V
		f func(index int, item V)
	}
	type testCase[V any] struct {
		name string
		args args[V]
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
				f: func(index int, item int) { t.Logf("%d: %d", index, item) },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForEachI(tt.args.s, tt.args.f)
		})
	}
}

func TestReverse(t *testing.T) {
	type args[V any] struct {
		s []V
	}
	type testCase[V any] struct {
		name string
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				s: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("Reverse() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
