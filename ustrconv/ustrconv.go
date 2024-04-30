package ustrconv

import "strconv"

func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func UnsafeParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ParseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func UnsafeParseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func ParseUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func UnsafeParseUint64(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

func ParseHexInt(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}

func UnsafeParseHexInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 16, 64)
	return i
}

func ParseHexUint(s string) (uint64, error) {
	return strconv.ParseUint(s, 16, 64)
}

func UnsafeParseHexUint(s string) uint64 {
	i, _ := strconv.ParseUint(s, 16, 64)
	return i
}
