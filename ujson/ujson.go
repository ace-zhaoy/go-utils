package ujson

import (
	"encoding/json"
	"unsafe"
)

func ToJson(v any) string {
	b, _ := json.Marshal(v)
	return *(*string)(unsafe.Pointer(&b))
}
