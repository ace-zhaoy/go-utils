package umap

import (
	"reflect"
	"testing"
)

func TestFindKeyByValue_int(t *testing.T) {
	type args[K comparable, V comparable] struct {
		m map[K]V
		v V
	}
	type testCase[K comparable, V comparable] struct {
		name   string
		args   args[K, V]
		wantK  K
		wantOk bool
	}
	tests := []testCase[int, int]{
		{
			name: "test",
			args: args[int, int]{
				m: map[int]int{1: 1, 2: 2, 3: 3},
				v: 2,
			},
			wantK:  2,
			wantOk: true,
		},
		{
			name: "not exist",
			args: args[int, int]{
				m: map[int]int{1: 1, 2: 2, 3: 3},
				v: 4,
			},
			wantK:  0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotOk := FindKeyByValue(tt.args.m, tt.args.v)
			if !reflect.DeepEqual(gotK, tt.wantK) {
				t.Errorf("FindKeyByValue() gotK = %v, want %v", gotK, tt.wantK)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FindKeyByValue() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestFindKeyByValue_string(t *testing.T) {
	type args[K comparable, V comparable] struct {
		m map[K]V
		v V
	}
	type testCase[K comparable, V comparable] struct {
		name   string
		args   args[K, V]
		wantK  K
		wantOk bool
	}
	tests := []testCase[string, string]{
		{
			name: "test",
			args: args[string, string]{
				m: map[string]string{"a": "a", "b": "b", "c": "c"},
				v: "b",
			},
			wantK:  "b",
			wantOk: true,
		},
		{
			name: "not exist",
			args: args[string, string]{
				m: map[string]string{"a": "a", "b": "b", "c": "c"},
				v: "d",
			},
			wantK:  "",
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotOk := FindKeyByValue(tt.args.m, tt.args.v)
			if !reflect.DeepEqual(gotK, tt.wantK) {
				t.Errorf("FindKeyByValue() gotK = %v, want %v", gotK, tt.wantK)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FindKeyByValue() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestFindKeyByValue_ptr(t *testing.T) {
	type args[K comparable, V comparable] struct {
		m map[K]V
		v V
	}
	type testCase[K comparable, V comparable] struct {
		name   string
		args   args[K, V]
		wantK  K
		wantOk bool
	}
	type User struct {
		Name string
	}
	b := &User{Name: "b"}
	tests := []testCase[int, *User]{
		{
			name: "test",
			args: args[int, *User]{
				m: map[int]*User{1: {Name: "a"}, 2: b, 3: {Name: "c"}},
				v: b,
			},
			wantK:  2,
			wantOk: true,
		},
		{
			name: "not exist",
			args: args[int, *User]{
				m: map[int]*User{1: {Name: "a"}, 2: {Name: "b"}, 3: {Name: "c"}},
				v: &User{Name: "b"},
			},
			wantK:  0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotOk := FindKeyByValue(tt.args.m, tt.args.v)
			if !reflect.DeepEqual(gotK, tt.wantK) {
				t.Errorf("FindKeyByValue() gotK = %v, want %v", gotK, tt.wantK)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FindKeyByValue() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestForeach(t *testing.T) {
	testCases := []struct {
		name string
		m    map[int]string
		f    func(k int, v string)
	}{
		{
			name: "empty map",
			m:    map[int]string{},
			f:    func(k int, v string) {},
		},
		{
			name: "non-empty map",
			m:    map[int]string{1: "apple", 2: "banana", 3: "cherry"},
			f: func(k int, v string) {
				t.Logf("%d: %s", k, v)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Foreach(tc.m, tc.f)
		})
	}
}

func TestKeyExists(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
		k K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want bool
	}
	tests := []testCase[string, int]{
		{
			name: "exist",
			args: args[string, int]{
				m: map[string]int{"a": 1, "b": 2, "c": 3},
				k: "b",
			},
			want: true,
		},
		{
			name: "not exist",
			args: args[string, int]{
				m: map[string]int{"a": 1, "b": 2, "c": 3},
				k: "d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyExists(tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("KeyExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
