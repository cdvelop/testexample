package main

import "testing"

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		target    int
		expectativa int
	}{
		{"Caso 1: el target aparece 3 veces", []int{1, 2, 2, 3, 2}, 2, 3},
		{"Caso 2: el target no aparece", []int{1, 2, 3, 4, 5}, 6, 0},
		{"Caso 3: slice vacío", []int{}, 1, 0},
		{"Caso 4: el target aparece al inicio", []int{5, 1, 2, 3}, 5, 1},
		{"Caso 5: el target aparece al final", []int{1, 2, 3, 5}, 5, 1},
		{"Caso 6: todos los elementos son el target", []int{7, 7, 7, 7}, 7, 4},
		{"Caso 7: target negativo en slice de negativos", []int{-1, -2, -1, -3}, -1, 2},
		{"Caso 8: target cero", []int{0, 1, 0, 2, 0}, 0, 3},
		{"Caso 9: slice con un solo elemento que es el target", []int{10}, 10, 1},
		{"Caso 10: slice con un solo elemento que no es el target", []int{10}, 5, 0},
		{"Caso 11: target con números grandes", []int{100, 200, 100, 50}, 100, 2},
		{"Caso 12: números negativos y positivos", []int{-1, 1, -1, 1}, -1, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountOccurrences(tt.nums, tt.target)
			if got != tt.expectativa {
				t.Errorf("CountOccurrences(%v, %d) = %d; se esperaba %d", tt.nums, tt.target, got, tt.expectativa)
			}
		})
	}
}