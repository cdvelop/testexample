package main

import (
	"reflect"
	"testing"
)

func TestSortInts(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expectativa []int
	}{
		{"Caso 1: Slice desordenado", []int{3, 1, 2}, []int{1, 2, 3}},
		{"Caso 2: Slice vacío", []int{}, []int{}},
		{"Caso 3: Ya ordenado", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"Caso 4: Ordenado en reversa", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{"Caso 5: Con números negativos", []int{-5, 5, 0, -1}, []int{-5, -1, 0, 5}},
		{"Caso 6: Con duplicados", []int{9, 4, 2, 9, 2}, []int{2, 2, 4, 9, 9}},
		{"Caso 7: Un solo elemento", []int{100}, []int{100}},
		{"Caso 8: Todos los elementos iguales", []int{7, 7, 7}, []int{7, 7, 7}},
		{"Caso 9: Slice nulo", nil, []int{}},
		{"Caso 10: Verifica que el original no se modifique", []int{5, 4, 6}, []int{4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalInput := make([]int, len(tt.input))
			copy(originalInput, tt.input)

			got := SortInts(tt.input)
			if !reflect.DeepEqual(got, tt.expectativa) {
				t.Errorf("SortInts(%v) = %v; want %v", tt.input, got, tt.expectativa)
			}

			// Verificación especial para el caso 10
			if tt.name == "Verifica que el original no se modifique" {
				if !reflect.DeepEqual(tt.input, originalInput) {
					t.Errorf("El slice original fue modificado: got %v; want %v", tt.input, originalInput)
				}
			}
		})
	}
}
