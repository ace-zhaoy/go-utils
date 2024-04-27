package ureflect

import (
	"testing"
	"time"
)

func TestHasField(t *testing.T) {
	type Person struct {
		Name  string
		Age   int
		Email string
	}

	// Test when the struct has the field
	p1 := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	if !HasField(p1, "Name") {
		t.Errorf("Expected true, got false")
	}

	// Test when the struct does not have the field
	if HasField(p1, "Address") {
		t.Errorf("Expected false, got true")
	}

	// Test when the value is a pointer to a struct with the field
	p2 := &Person{Name: "Bob", Age: 25, Email: "bob@example.com"}
	if !HasField(*p2, "Name") {
		t.Errorf("Expected true, got false")
	}

	// Test when the value is a pointer to a struct without the field
	if HasField(p2, "Address") {
		t.Errorf("Expected false, got true")
	}

	// Test when the value is a nil pointer
	var p3 *Person
	if !HasField(p3, "Name") {
		t.Errorf("Expected true, got false")
	}

	// Test when the value is not a struct or a pointer to a struct
	var i int
	if HasField(i, "Name") {
		t.Errorf("Expected false, got true")
	}
}

func TestIsZero(t *testing.T) {
	var f func()
	tests := []struct {
		name  string
		input any
		want  bool
	}{
		// Basic types and their zero values
		{"Zero int", 0, true},
		{"Non-zero int", 1, false},
		{"Zero float", 0.0, true},
		{"Non-zero float", 0.1, false},
		{"Zero string", "", true},
		{"Non-zero string", "a", false},
		{"Zero bool", false, true},
		{"Non-zero bool", true, false},

		// Pointers to both zero and non-zero values
		{"Nil pointer", (*int)(nil), true},
		{"Pointer to zero int", new(int), false},
		{"Pointer to non-zero int", &[]int{1}[0], false},

		// Slices, which are nil by default
		{"Nil slice", []int(nil), true},
		{"Empty slice", []int{}, false},
		{"Non-empty slice", []int{0}, false},

		// Arrays which can't be nil, but can be zero-valued
		{"Zero-valued array", [1]int{}, true},
		{"Non-zero-valued array", [1]int{1}, false},

		// Structs, which are never nil, but can be zero-valued
		{"Zero-valued struct", struct{ a int }{}, true},
		{"Non-zero-valued struct", struct{ a int }{1}, false},

		// Custom types implementing Zeroable
		{"Zero time", time.Time{}, true},
		{"Non-zero time", time.Now(), false},

		// Maps, which can be nil
		{"Nil map", map[string]int(nil), true},
		{"Empty map", map[string]int{}, false},
		{"Non-empty map", map[string]int{"a": 0}, false},

		// Channels, which can be nil
		{"Nil channel", (chan int)(nil), true},
		{"Non-nil channel", make(chan int), false},

		{"nil", nil, true},
		{"func", f, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.input); got != tt.want {
				t.Errorf("IsZero(%#v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsRealZero(t *testing.T) {
	type MyInt int
	var f func()
	tests := []struct {
		name string
		val  any
		want bool
	}{
		{"IntZero", 0, true},
		{"IntNonZero", 1, false},
		{"FloatZero", 0.0, true},
		{"FloatNonZero", 1.1, false},
		{"StringEmpty", "", true},
		{"StringNonEmpty", "notEmpty", false},
		{"BoolFalse", false, true},
		{"BoolTrue", true, false},
		{"PointerNil", (*int)(nil), true},
		{"PointerNonNil", new(int), true}, // zero value for int is 0
		{"SliceNil", []int(nil), true},
		{"SliceEmpty", []int{}, true},
		{"SliceNonEmpty", []int{0}, false},
		{"MapNil", map[int]any(nil), true},
		{"MapEmpty", map[int]any{}, true},
		{"MapNonEmpty", map[int]any{1: 1}, false},
		{"InterfaceNil", interface{}(nil), true},
		{"InterfaceNonNil", interface{}(0), true},
		{"CustomTypeZero", MyInt(0), true},
		{"CustomTypeNonZero", MyInt(1), false},
		{"Zero time", time.Time{}, true},
		{"funcNil", f, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRealZero(tt.val); got != tt.want {
				t.Errorf("IsRealZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
