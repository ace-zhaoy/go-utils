package ucompare

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a    T
		args []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Max",
			args: args[int]{
				args: []int{1, 2, 5, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a    T
		args []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Min",
			args: args[int]{
				args: []int{5, 2, 1, 4},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
