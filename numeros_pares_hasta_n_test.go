package main

import (
	"reflect"
	"testing"
)

func TestNumerosParesHastaN(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		expectativa []int
	}{
		{"Caso 1: Hasta un número par", 6, []int{0, 2, 4, 6}},
		{"Caso 2: Hasta un número impar", 7, []int{0, 2, 4, 6}},
		{"Caso 3: Hasta 0", 0, []int{0}},
		{"Caso 4: Hasta 1", 1, []int{0}},
		{"Caso 5: Hasta 10", 10, []int{0, 2, 4, 6, 8, 10}},
		{"Caso 6: Input negativo", -5, []int{}},
		{"Caso 7: Hasta 2", 2, []int{0, 2}},
		{"Caso 8: Hasta 3", 3, []int{0, 2}},
		{"Caso 9: Hasta 4", 4, []int{0, 2, 4}},
		{"Caso 10: Límite superior grande", 14, []int{0, 2, 4, 6, 8, 10, 12, 14}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumerosParesHastaN(tt.n)
			if !reflect.DeepEqual(got, tt.expectativa) {
				t.Errorf("EvenNumbersUntil(%d) = %v; want %v", tt.n, got, tt.expectativa)
			}
		})
	}
}
