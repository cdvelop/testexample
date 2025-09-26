package main

import (
	"reflect"
	"testing"
)

func TestMultiplicationTable(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		expectativa []int
	}{
		{"Caso 1: Tabla del 3", 3, []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}},
		{"Caso 2: Tabla del 1", 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Caso 3: Tabla del 10", 10, []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}},
		{"Caso 4: Tabla del 0", 0, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"Caso 5: Tabla de un número negativo", -2, []int{-2, -4, -6, -8, -10, -12, -14, -16, -18, -20}},
		{"Caso 6: Tabla del 5", 5, []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}},
		{"Caso 7: Tabla del 7", 7, []int{7, 14, 21, 28, 35, 42, 49, 56, 63, 70}},
		{"Caso 8: Tabla del 9", 9, []int{9, 18, 27, 36, 45, 54, 63, 72, 81, 90}},
		{"Caso 9: Tabla de un número grande", 100, []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}},
		{"Caso 10: Tabla del 2", 2, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MultiplicationTable(tt.n)
			if !reflect.DeepEqual(got, tt.expectativa) {
				t.Errorf("MultiplicationTable(%d) = %v; want %v", tt.n, got, tt.expectativa)
			}
		})
	}
}
