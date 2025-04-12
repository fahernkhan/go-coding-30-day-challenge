// tests/dp_test.go
package tests

import (
	"main"
	"testing"
)

func TestGetTotalX(t *testing.T) {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}
	expected := int32(3)

	result := main.GetTotalX(a, b)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
