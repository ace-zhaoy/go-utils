package umap

import (
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func Foreach[K comparable, V any](m map[K]V, f func(k K, v V)) {
	for k, v := range m {
		f(k, v)
	}
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
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

func ToStructByTag(input any, output any, tagName string) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   output,
		TagName:  tagName,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func ToStruct(input any, output any) error {
	return ToStructByTag(input, output, "json")
}

func ParseByTag(input any, tagName string) map[string]any {
	s := structs.New(input)
	s.TagName = tagName
	return s.Map()
}

func Parse(input any) map[string]any {
	return ParseByTag(input, "json")
}
