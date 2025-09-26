package main // Paquete principal

import "testing" // Para pruebas unitarias

func TestCountOccurrences(t *testing.T) { // Test para CountOccurrences
	tests := []struct {
		name        string // Nombre del caso de prueba
		nums        []int  // Slice de entrada
		target      int    // Valor a buscar
		expectativa int    // Resultado esperado
	}{
		{"Caso 1: el target aparece 3 veces", []int{1, 2, 2, 3, 2}, 2, 3},
		{"Caso 2: el target no aparece", []int{1, 2, 3, 4, 5}, 6, 0},
		{"Caso 3: slice vacío", []int{}, 1, 0},
		{"Caso 4: el target aparece al inicio", []int{5, 1, 2, 3}, 5, 1},
		{"Caso 5: el target aparece al final", []int{1, 2, 3, 5}, 5, 1},
		{"Caso 6: todos los elementos son el target", []int{7, 7, 7, 7}, 7, 4},
		{"Caso 7: target negativo en slice de negativos", []int{-1, -2, -1, -3}, -1, 2},
		{"Caso 8: target cero", []int{0, 1, 0, 2, 0}, 0, 3},
		{"Caso 9: slice con un solo elemento que es el target", []int{10}, 10, 1},
		{"Caso 10: slice con un solo elemento que no es el target", []int{10}, 5, 0},
		{"Caso 11: target con números grandes", []int{100, 200, 100, 50}, 100, 2},
		{"Caso 12: números negativos y positivos", []int{-1, 1, -1, 1}, -1, 2},
	}

	for _, tt := range tests { // Itera sobre cada caso de prueba
		t.Run(tt.name, func(t *testing.T) { // Ejecuta subtest
			got := CountOccurrences(tt.nums, tt.target) // Llama a la función
			if got != tt.expectativa {                  // Compara resultado
				t.Errorf("CountOccurrences(%v, %d) = %d; se esperaba %d", tt.nums, tt.target, got, tt.expectativa) // Falla si no coincide
			}
		})
	}
}
