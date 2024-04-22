package ustring

import (
	"reflect"
	"strings"
	"unsafe"
)

// Bytes2Str converts []byte to string,no-copy
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Str2Bytes converts string to []byte,no-copy
func Str2Bytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

func Split(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}
