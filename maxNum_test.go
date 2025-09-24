package testexample

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{name: "caso1", a: 3, b: 7, expected: 7},
		{name: "caso2", a: 10, b: 5, expected: 10},
		{name: "caso3", a: -2, b: -5, expected: -2},
		{name: "caso4", a: 0, b: 0, expected: 0},
		{name: "caso5", a: 9, b: 8, expected: 9},
		{name: "caso6", a: -1, b: 1, expected: 1},
		{name: "caso7", a: 100, b: 1000, expected: 1000},
		{name: "caso8", a: -10, b: -1, expected: -1},
		{name: "caso9", a: 321, b: 123, expected: 321},
		{name: "caso10", a: 41, b: 42, expected: 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Max(%v, %v) = %v; expected %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
