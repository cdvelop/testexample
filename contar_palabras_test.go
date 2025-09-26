package main

import "testing"

func TestContarPalabras(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectativa int
	}{
		{"Caso 1: Frase de ejemplo", "hola mundo go", 3},
		{"Caso 2: String vacío", "", 0},
		{"Caso 3: Una sola palabra", "gopher", 1},
		{"Caso 4: Con espacios extra al inicio y final", "  hola mundo  ", 2},
		{"Caso 5: Con múltiples espacios entre palabras", "uno   dos    tres", 3},
		{"Caso 6: String con solo espacios", "   ", 0},
		{"Caso 7: Con tabuladores y saltos de línea", "linea1\nlinea2\tlinea3", 3},
		{"Caso 8: Frase larga", "este es un certamen de testing en go", 8},
		{"Caso 9: Con números y palabras", "el 1 es un numero", 5},
		{"Caso 10: Con puntuación pegada", "hola,mundo", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContarPalabras(tt.input)
			if got != tt.expectativa {
				t.Errorf("ContarPalabras(%q) = %d; want %d", tt.input, got, tt.expectativa)
			}
		})
	}
}
