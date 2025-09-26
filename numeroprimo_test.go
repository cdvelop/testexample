package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		expectativa bool
	}{
		{"Caso 1: 2 es primo", 2, true},
		{"Caso 2: 7 es primo", 7, true},
		{"Caso 3: 1 (no es primo)", 1, false},
		{"Caso 4: 4 (no es primo)", 4, false},
		{"Caso 5: 9 (no es primo)", 9, false},
		{"Caso 6: 10 (no es primo)", 10, false},
		{"Caso 7: 13 es primo", 13, true},
		{"Caso 8: 0 (no es primo)", 0, false},
		{"Caso 9: 25 (no es primo)", 25, false},
		{"Caso 10: 97 es primo", 97, true},
		{"Caso 11: -3 (no es primo)", -3, false},
		{"Caso 12: 121 (no es primo)", 121, false}, // 11 * 11
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPrime(tt.n)
			if got != tt.expectativa {
				t.Errorf("IsPrime(%d) = %t; se esperaba %t", tt.n, got, tt.expectativa)
			}
		})
	}
}