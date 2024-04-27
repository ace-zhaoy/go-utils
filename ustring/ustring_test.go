package ustring

import (
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

func TestBytes2Str(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{b: []byte("test")},
			want: "test",
		},
		{
			name: "你好",
			args: args{b: []byte("你好")},
			want: "你好",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2Str(tt.args.b); got != tt.want {
				t.Errorf("Bytes2Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes2StrNoCopy(t *testing.T) {
	b := []byte("test")
	s := Bytes2Str(b)

	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	if bh.Data != sh.Data {
		t.Errorf("Bytes2Str copied memory, bh.Data = %p, sh.Data = %p", unsafe.Pointer(bh.Data), unsafe.Pointer(sh.Data))
	}

	runtime.GC()
	if bh.Data != sh.Data {
		t.Errorf("Bytes2Str copied memory after GC, bh.Data = %p, sh.Data = %p", unsafe.Pointer(bh.Data), unsafe.Pointer(sh.Data))
	}
}

func TestStr2Bytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantB []byte
	}{
		{"hello", args{"hello"}, []byte("hello")},
		{"你好", args{"你好"}, []byte("你好")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := Str2Bytes(tt.args.s); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("Str2Bytes() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestStr2BytesNoCopy(t *testing.T) {
	s := "hello"
	b := Str2Bytes(s)

	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	if bh.Data != sh.Data {
		t.Errorf("Str2Bytes copied memory, bh.Data = %p, sh.Data = %p", unsafe.Pointer(bh.Data), unsafe.Pointer(sh.Data))
	}

	runtime.GC()
	if bh.Data != sh.Data {
		t.Errorf("Str2Bytes copied memory after GC, bh.Data = %p, sh.Data = %p", unsafe.Pointer(bh.Data), unsafe.Pointer(sh.Data))
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{s: "a,b,c", sep: ","},
			want: []string{"a", "b", "c"},
		},
		{
			name: "empty",
			args: args{s: "", sep: ","},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrTrim(t *testing.T) {
	type args struct {
		s      string
		cutset []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{s: " 你好  ", cutset: []string{" "}},
			want: "你好",
		},
		{
			name: "empty",
			args: args{s: "", cutset: []string{" "}},
			want: "",
		},
		{
			name: "cutset nil",
			args: args{s: " 你好\n\t  ", cutset: nil},
			want: "你好",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrTrim(tt.args.s, tt.args.cutset...); got != tt.want {
				t.Errorf("StrTrim() = %v, want %v", got, tt.want)
			}
		})
	}
}
