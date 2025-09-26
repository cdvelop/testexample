package main // Paquete principal

import "testing" // Para pruebas unitarias

func TestLongestWord(t *testing.T) { // Test para LongestWord
	tests := []struct {
		name        string // Nombre del caso de prueba
		s           string // Cadena de entrada
		expectativa string // Resultado esperado
	}{
		{"Caso 1: ejemplo del certamen", "go es divertido", "divertido"},
		{"Caso 2: cadena vacía", "", ""},
		{"Caso 3: solo espacios", "   ", ""},
		{"Caso 4: una sola palabra", "programacion", "programacion"},
		{"Caso 5: empate de longitud (primera encontrada)", "largo uno largo dos", "largo"},
		{"Caso 6: espacios al inicio y final", "  go es un lenguaje ", "lenguaje"},
		{"Caso 7: números en la cadena", "1234 1234567 12", "1234567"},
		{"Caso 8: multiples palabras de la misma longitud", "coche barco avion", "coche"},
		{"Caso 9: palabra al final es la mas larga", "la casa azul grande", "grande"},
		{"Caso 10: multiples palabras de la misma longitud en desorden", "coche avion barco", "coche"},
	}

	for _, tt := range tests { // Itera sobre cada caso de prueba
		t.Run(tt.name, func(t *testing.T) { // Ejecuta subtest
			got := LongestWord(tt.s)   // Llama a la función
			if got != tt.expectativa { // Compara resultado
				t.Errorf("LongestWord(%q) = %q; se esperaba %q", tt.s, got, tt.expectativa) // Falla si no coincide
			}
		})
	}
}
