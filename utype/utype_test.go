package utype

import (
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testInstance := testStruct{Name: "John", Age: 30}

	ptr := Pointer(testInstance)

	if ptr == nil {
		t.Errorf("Pointer should not be nil")
	}

	if *ptr != testInstance {
		t.Errorf("Pointer value does not match the original struct")
	}
}

func TestDereference(t *testing.T) {
	type args[T any] struct {
		p            *T
		defaultValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				p:            Pointer(5),
				defaultValue: 10,
			},
			want: 5,
		},
		{
			name: "test2",
			args: args[int]{
				p:            nil,
				defaultValue: 10,
			},
			want: 10,
		},
		{
			name: "test3",
			args: args[int]{
				p:            Pointer(0),
				defaultValue: 10,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dereference(tt.args.p, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dereference() = %v, want %v", got, tt.want)
			}
		})
	}
}
