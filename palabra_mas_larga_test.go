package main

import "testing"

func TestLongestWord(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectativa string
	}{
		{"Caso 1: Ejemplo básico", "go es divertido", "divertido"},
		{"Caso 2: String vacío", "", ""},
		{"Caso 3: Empate, retorna la primera", "un gran dia", "gran"},
		{"Caso 4: Una sola palabra", "otorrinolaringologo", "otorrinolaringologo"},
		{"Caso 5: Con puntuación", "esto, es. una_prueba", "una_prueba"},
		{"Caso 6: String con solo espacios", "   ", ""},
		{"Caso 7: La más larga al inicio", "esternocleidomastoideo es un musculo", "esternocleidomastoideo"},
		{"Caso 8: Todas las palabras de igual longitud", "el la lo se", "el"},
		{"Caso 9: Con números", "1 12 123 1234", "1234"},
		{"Caso 10: Frase compleja", "bienvenidos al certamen de programacion", "programacion"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestWord(tt.input)
			if got != tt.expectativa {
				t.Errorf("LongestWord(%q) = %q; want %q", tt.input, got, tt.expectativa)
			}
		})
	}
}
