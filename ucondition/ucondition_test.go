package ucondition

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type args[T any] struct {
		condition  bool
		trueValue  T
		falseValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[any]{
		{
			name: "test 1",
			args: args[any]{condition: true, trueValue: 1, falseValue: 0},
			want: 1,
		},
		{
			name: "test 0",
			args: args[any]{condition: false, trueValue: 1, falseValue: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.condition, tt.args.trueValue, tt.args.falseValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfF(t *testing.T) {
	type args[T any] struct {
		condition bool
		trueFunc  func() T
		falseFunc func() T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}

	tests := []testCase[any]{
		{
			name: "test 1",
			args: args[any]{condition: true, trueFunc: func() any { return 1 }, falseFunc: func() any { return 0 }},
			want: 1,
		},
		{
			name: "test 0",
			args: args[any]{condition: false, trueFunc: func() any { return 1 }, falseFunc: func() any { return 0 }},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfF(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfFF(t *testing.T) {
	type User struct {
		Name string
	}
	var t1 *User
	t2 := &User{
		Name: "John",
	}
	type args[T any] struct {
		condition bool
		trueValue T
		falseFunc func() T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[any]{
		{
			name: "test 1",
			args: args[any]{condition: t1 == nil, trueValue: "", falseFunc: func() any { return t1.Name }},
			want: "",
		},
		{
			name: "test 2",
			args: args[any]{condition: t2 == nil, trueValue: "", falseFunc: func() any { return t2.Name }},
			want: "John",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfFF(tt.args.condition, tt.args.trueValue, tt.args.falseFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfFF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfTF(t *testing.T) {
	type User struct {
		Name string
	}
	var t1 *User
	t2 := &User{
		Name: "John",
	}
	type args[T any] struct {
		condition  bool
		trueFunc   func() T
		falseValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[any]{
		{
			name: "test 1",
			args: args[any]{condition: t1 != nil, trueFunc: func() any { return t1.Name }, falseValue: ""},
			want: "",
		},
		{
			name: "test 2",
			args: args[any]{condition: t2 != nil, trueFunc: func() any { return t2.Name }, falseValue: ""},
			want: "John",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfTF(tt.args.condition, tt.args.trueFunc, tt.args.falseValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfTF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfZero_int(t *testing.T) {
	type args[T any] struct {
		value        T
		defaultValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "ZeroValueWithDefault",
			args: args[int]{value: 0, defaultValue: 10},
			want: 10,
		},
		{
			name: "NonZeroValue",
			args: args[int]{value: 5, defaultValue: 10},
			want: 5,
		},
		{
			name: "NegativeValue",
			args: args[int]{value: -5, defaultValue: 10},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfZero(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("IfZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfZero_int_pointer(t *testing.T) {
	type args[T any] struct {
		value        T
		defaultValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	var (
		five = 5
		zero = 0
		ten  = 10
	)
	tests := []testCase[*int]{
		{
			name: "PointerToZeroValueWithDefault",
			args: args[*int]{value: &zero, defaultValue: &ten},
			want: &zero,
		},
		{
			name: "PointerToNonZeroValue",
			args: args[*int]{value: &five, defaultValue: &ten},
			want: &five,
		},
		{
			name: "NilValueWithDefault",
			args: args[*int]{value: nil, defaultValue: &ten},
			want: &ten,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfZero(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("IfZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue(t *testing.T) {
	type args[T any] struct {
		value        T
		defaultValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	var (
		five = 5
		zero = 0
		ten  = 10
	)
	tests := []testCase[*int]{
		{
			name: "case1",
			args: args[*int]{value: &zero, defaultValue: &ten},
			want: &zero,
		},
		{
			name: "case2",
			args: args[*int]{value: &five, defaultValue: &ten},
			want: &five,
		},
		{
			name: "case3",
			args: args[*int]{value: nil, defaultValue: &ten},
			want: &ten,
		},
		{
			name: "case4",
			args: args[*int]{value: &ten, defaultValue: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Value(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
