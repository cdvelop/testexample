package testexample

import "testing"

func TestMaxNum(t *testing.T) {
	tests := []struct {
		name        string // nombre del caso
		num1, num2  int    // datos de entrada
		expectativa int    // expectativa
	}{
		{name: "Caso 1", num1: 5, num2: 3, expectativa: 5},
		{name: "Caso 2", num1: 0, num2: 5, expectativa: 5},
		{name: "Caso 3", num1: 1, num2: 2, expectativa: 2},
		{name: "Caso 4", num1: 2, num2: 11, expectativa: 11},
		{name: "Caso 5", num1: 3, num2: 4, expectativa: 4},
		{name: "Caso 6", num1: 4, num2: 1, expectativa: 4},
		{name: "Caso 7", num1: 6, num2: -3, expectativa: 6},
		{name: "Caso 8", num1: 7, num2: 4, expectativa: 7},
		{name: "Caso 9", num1: 8, num2: 5551, expectativa: 5551},
		{name: "Caso 10", num1: 9, num2: 4, expectativa: 9},

		// TODO: agregar al menos 10 casos
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumMax(tt.num1, tt.num2)
			if got != tt.expectativa {
				t.Errorf("NumMax(%d, %d) = %d; want %d", tt.num1, tt.num2, got, tt.expectativa)
			}
		})
	}
}
