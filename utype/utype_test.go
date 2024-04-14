package utype

import (
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
