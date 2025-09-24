package testexample

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Factorial de 0", 0, 1},
		{"Factorial de 1", 1, 1},
		{"Factorial de 2", 2, 2},
		{"Factorial de 3", 3, 6},
		{"Factorial de 5", 5, 120},
		{"Factorial de 6", 6, 720},
		{"Factorial de 10", 10, 3628800},
		{"Factorial de 12", 12, 479001600},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Factorial(tt.input)
			if got != tt.expected {
				t.Errorf("Factorial(%d) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}
