package testexample

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name        string // nombre del caso
		word        string // datos de entrada
		expectativa bool   // expectativa
	}{
		{name: "Caso 1", word: "paralelepipedo", expectativa: false},
		{name: "Caso 2", word: "radar", expectativa: true},
		{name: "Caso 3", word: "camello", expectativa: false},
		{name: "Caso 4", word: "anilina", expectativa: true},
		{name: "Caso 5", word: "amoma", expectativa: true},
		{name: "Caso 6", word: "pero", expectativa: false},
		{name: "Caso 7", word: "licuado", expectativa: false},
		{name: "Caso 8", word: "reconocer", expectativa: true},
		{name: "Caso 9", word: "aaaaaaa", expectativa: true},
		{name: "Caso 10", word: "compadre", expectativa: false},

		// TODO: agregar al menos 10 casos
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.word)
			if got != tt.expectativa {
				t.Errorf("Error en %s: se esperaba %v y se obtuvo %v", tt.name, tt.expectativa, got)
			}
		})
	}
}
