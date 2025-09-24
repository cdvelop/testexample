package testexample

import "testing"

func TestMax(t *testing.T) {
	a := 42
	b := 27
	expected := 42

	result := max(a, b)
	if result != expected {
		t.Errorf("Expected max(%d, %d) to be %d but got %d", a, b, expected, result)
	}
}
