package main

import "testing"

func TestContarOcurrencias(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		target      int
		expectativa int
	}{
		{"Caso 1: Ejemplo básico", []int{1, 2, 2, 3, 2}, 2, 3},
		{"Caso 2: Target no presente", []int{1, 2, 3}, 4, 0},
		{"Caso 3: Slice vacío", []int{}, 5, 0},
		{"Caso 4: Todos son el target", []int{8, 8, 8}, 8, 3},
		{"Caso 5: Con números negativos", []int{-1, 2, -1, 3, -1}, -1, 3},
		{"Caso 6: Target es cero", []int{0, 1, 0, 2, 0}, 0, 3},
		{"Caso 7: Target al inicio y final", []int{9, 1, 2, 9}, 9, 2},
		{"Caso 8: Un solo elemento (match)", []int{7}, 7, 1},
		{"Caso 9: Un solo elemento (no match)", []int{7}, 8, 0},
		{"Caso 10: Slice nulo", nil, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContarOcurrencias(tt.nums, tt.target)
			if got != tt.expectativa {
				t.Errorf("ContarOcurrencias(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.expectativa)
			}
		})
	}
}
