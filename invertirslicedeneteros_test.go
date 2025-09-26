package main

import (
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		expectativa []int
	}{
		// Casos de prueba para la función ReverseSlice.
		{"Caso 1: ejemplo del certamen", []int{1, 2, 3}, []int{3, 2, 1}},
		{"Caso 2: slice vacío", []int{}, []int{}},
		{"Caso 3: un solo elemento", []int{10}, []int{10}},
		{"Caso 4: slice con números pares", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"Caso 5: slice ya invertido", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"Caso 6: números negativos", []int{-1, -2, -3}, []int{-3, -2, -1}},
		{"Caso 7: combinación de positivos y negativos", []int{-1, 0, 1}, []int{1, 0, -1}},
		{"Caso 8: valores repetidos", []int{1, 2, 1, 2}, []int{2, 1, 2, 1}},
		{"Caso 9: slice de longitud 5", []int{5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5}},
		{"Caso 10: slice de longitud 6", []int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{"Caso 11: slice con números grandes", []int{100, 2000, 300}, []int{300, 2000, 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseSlice(tt.nums)
			if !reflect.DeepEqual(got, tt.expectativa) {
				t.Errorf("ReverseSlice(%v) = %v; se esperaba %v", tt.nums, got, tt.expectativa)
			}
		})
	}
}