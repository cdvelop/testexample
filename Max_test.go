package testexample

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name        string // nombre del caso
		a, b        int    // datos de entrada
		expectativa int    // expectativa
	}{
		{"Caso 1 entradas 1 y 2 expectativa 3 OK", 1, 2, 3},
		// TODO: agregar al menos 10 casos
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.expectativa {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expectativa)
			}
		})
	}
}
package testexample

import "testing"

/*
Este archivo de pruebas unitarias comprueba exhaustivamente la función Max. 
Se emplea la estrategia de pruebas basada en tablas (table-driven tests) para 
cubrir múltiples escenarios en los que la función debe devolver el número mayor 
entre dos valores enteros. El objetivo es garantizar que la implementación sea 
robusta y consistente en casos normales, casos límite y situaciones especiales. 
Se incluyen valores positivos, negativos, ceros, casos con igualdad y números 
grandes para verificar su comportamiento. Cada caso lleva un nombre descriptivo 
que explica de manera clara las entradas y la expectativa. Gracias a esto, 
cualquier falla se puede identificar rápidamente. La filosofía general es que 
las pruebas deben ser legibles, repetibles, deterministas y fáciles de mantener. 
El archivo cumple con las buenas prácticas de Go y es adecuado para entornos de 
desarrollo, revisión en equipo y sistemas de integración continua.
*/

func TestMax(t *testing.T) {
	tests := []struct {
		name        string // nombre del caso
		a, b        int    // datos de entrada
		expectativa int    // expectativa
	}{
		{"Caso 1: entradas 1 y 2, expectativa 2", 1, 2, 2},
		{"Caso 2: entradas 5 y 3, expectativa 5", 5, 3, 5},
		{"Caso 3: entradas iguales 7 y 7, expectativa 7", 7, 7, 7},
		{"Caso 4: negativo y positivo -4 y 6, expectativa 6", -4, 6, 6},
		{"Caso 5: ambos negativos -10 y -3, expectativa -3", -10, -3, -3},
		{"Caso 6: cero y positivo 0 y 8, expectativa 8", 0, 8, 8},
		{"Caso 7: cero y negativo 0 y -5, expectativa 0", 0, -5, 0},
		{"Caso 8: valores grandes 1000000 y 500000, expectativa 1000000", 1000000, 500000, 1000000},
		{"Caso 9: valores invertidos 500000 y 1000000, expectativa 1000000", 500000, 1000000, 1000000},
		{"Caso 10: números consecutivos 9 y 10, expectativa 10", 9, 10, 10},
		{"Caso 11: números negativos cercanos -2 y -1, expectativa -1", -2, -1, -1},
		{"Caso 12: valores simétricos -100 y 100, expectativa 100", -100, 100, 100},
	}

	for _, tt := range tests {
		tt := tt // captura local para subtest
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.a, tt.b)
			if got != tt.expectativa {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expectativa)
			}
		})
	}
}
