package main

import "testing"

func TestNumeroPrimo(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		expectativa bool
	}{
		{"Caso 1: Primo pequeño", 7, true},
		{"Caso 2: No primo", 10, false},
		{"Caso 3: El número 2", 2, true},
		{"Caso 4: El número 3", 3, true},
		{"Caso 5: El número 4 (no primo)", 4, false},
		{"Caso 6: El número 1 (no primo por definición)", 1, false},
		{"Caso 7: El número 0 (no primo)", 0, false},
		{"Caso 8: Número negativo", -5, false},
		{"Caso 9: Primo grande", 97, true},
		{"Caso 10: No primo grande (cuadrado de 7)", 49, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumeroPrimo(tt.n)
			if got != tt.expectativa {
				t.Errorf("NumeroPrimo(%d) = %t; want %t", tt.n, got, tt.expectativa)
			}
		})
	}
}
