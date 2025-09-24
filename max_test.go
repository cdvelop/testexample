package main

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name        string // nombre del caso
		a, b        int    // datos de entrada
		expectativa int    // expectativa
	}{
		{"Caso 1: a mayor que b positivos", 7, 3, 7},
		{"Caso 2: b mayor que a positivos", 5, 10, 10},
		{"Caso 3: números iguales positivos", 8, 8, 8},
		{"Caso 4: a mayor que b negativos", -2, -5, -2},
		{"Caso 5: b mayor que a negativos", -8, -3, -3},
		{"Caso 6: números iguales negativos", -4, -4, -4},
		{"Caso 7: a positivo b negativo", 6, -2, 6},
		{"Caso 8: a negativo b positivo", -3, 9, 9},
		{"Caso 9: uno es cero a mayor", 5, 0, 5},
		{"Caso 10: uno es cero b mayor", 0, 7, 7},
		{"Caso 11: ambos son cero", 0, 0, 0},
		{"Caso 12: números grandes positivos", 1000, 999, 1000},
		{"Caso 13: números grandes negativos", -1000, -999, -999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.expectativa {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expectativa)
			}
		})
	}
}
