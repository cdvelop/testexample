package main

import "testing"

func TestMaximoEnSlice(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expectativa int
	}{
		{"Caso 1: Ejemplo básico", []int{3, 7, 2}, 7},
		{"Caso 2: Slice vacío", []int{}, 0},
		{"Caso 3: Con números negativos", []int{-10, -5, -20}, -5},
		{"Caso 4: Mixto de positivos y negativos", []int{-5, 10, 0, 4}, 10},
		{"Caso 5: Máximo al inicio", []int{100, 20, 30}, 100},
		{"Caso 6: Máximo al final", []int{20, 30, 100}, 100},
		{"Caso 7: Todos los elementos iguales", []int{8, 8, 8}, 8},
		{"Caso 8: Un solo elemento", []int{42}, 42},
		{"Caso 9: El máximo es 0", []int{-1, -5, 0}, 0},
		{"Caso 10: Slice nulo", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaximoEnSlice(tt.input)
			if got != tt.expectativa {
				t.Errorf("MaximoEnSlice(%v) = %d; want %d", tt.input, got, tt.expectativa)
			}
		})
	}
}
