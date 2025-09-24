package testexample

import (
	"testing"
)

func TestSumPositives(t *testing.T) {
	input := []int{1, -2, 7}
	expected := 8

	result := SumPositives(input)
	if result != expected {
		t.Errorf("expected %d, but got %d", expected, result)
	}
}
