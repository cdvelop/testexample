package main

import "testing"

func TestWordCount(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		expectativa int
	}{
		{"Caso 1: cadena simple", "hola mundo go", 3},
		{"Caso 2: cadena vacía", "", 0},
		{"Caso 3: solo espacios", "   ", 0},
		{"Caso 4: una sola palabra", "palabra", 1},
		{"Caso 5: espacios al inicio y final", "  hola mundo  ", 2},
		{"Caso 6: múltiples espacios entre palabras", "uno   dos    tres", 3},
		{"Caso 7: cadena con puntuación", "hola, mundo.", 2}, // strings.Fields maneja la puntuación como parte de la palabra
		{"Caso 8: números en la cadena", "1 2 3 4 5", 5},
		{"Caso 9: cadena larga", "este es un ejemplo de una cadena de texto mas larga para probar la funcion", 15},
		{"Caso 10: cadena con saltos de linea", "linea\ncon\nsaltos", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordCount(tt.s)
			if got != tt.expectativa {
				t.Errorf("WordCount(%q) = %d; se esperaba %d", tt.s, got, tt.expectativa)
			}
		})
	}
}