package testexample

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expectativa int
	}{
		{"Caso 1: números positivos pequeños", 2, 3, 5},
		{"Caso 2: suma con cero (a)", 0, 5, 5},
		{"Caso 3: suma con cero (b)", 7, 0, 7},
		{"Caso 4: ambos números cero", 0, 0, 0},
		{"Caso 5: números negativos", -3, -2, -5},
		{"Caso 6: positivo y negativo (suma positiva)", 10, -3, 7},
		{"Caso 7: positivo y negativo (suma negativa)", 5, -8, -3},
		{"Caso 8: positivo y negativo (suma cero)", 5, -5, 0},
		{"Caso 9: números grandes positivos", 1000000, 2000000, 3000000},
		{"Caso 10: números grandes negativos", -1500000, -500000, -2000000},
		{"Caso 11: máximo entero positivo pequeño", 100, 200, 300},
		{"Caso 12: combinación de extremos", -999, 1000, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expectativa {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expectativa)
			}
		})
	}
}
