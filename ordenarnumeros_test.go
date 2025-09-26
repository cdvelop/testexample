package main // Paquete principal

import (
	"reflect" // Para comparar slices
	"testing" // Para pruebas unitarias
)

func TestSortInts(t *testing.T) { // Test para SortInts
	tests := []struct {
		name        string // Nombre del caso de prueba
		nums        []int  // Slice de entrada
		expectativa []int  // Resultado esperado
	}{
		// Casos de prueba para la función SortInts.
		{"Caso 1: slice desordenado", []int{3, 1, 2}, []int{1, 2, 3}},
		{"Caso 2: slice ya ordenado", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Caso 3: slice en orden inverso", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"Caso 4: slice vacío", []int{}, []int{}},
		{"Caso 5: slice con un solo elemento", []int{10}, []int{10}},
		{"Caso 6: números repetidos", []int{5, 2, 5, 1, 2}, []int{1, 2, 2, 5, 5}},
		{"Caso 7: números negativos", []int{-3, 0, -1, 2}, []int{-3, -1, 0, 2}},
		{"Caso 8: combinación de positivos y negativos", []int{1, -5, 3, -2, 0}, []int{-5, -2, 0, 1, 3}},
		{"Caso 9: slice con números grandes", []int{999, 10, 50, 200}, []int{10, 50, 200, 999}},
		{"Caso 10: slice con cero y números positivos", []int{5, 0, 3, 1}, []int{0, 1, 3, 5}},
	}

	for _, tt := range tests { // Itera sobre cada caso de prueba
		t.Run(tt.name, func(t *testing.T) { // Ejecuta subtest
			got := SortInts(tt.nums) // Llama a la función
			// reflect.DeepEqual se usa para comparar el contenido de los slices, ya que '==' no funciona para ellos.
			if !reflect.DeepEqual(got, tt.expectativa) { // Compara resultado
				t.Errorf("SortInts(%v) = %v; se esperaba %v", tt.nums, got, tt.expectativa) // Falla si no coincide
			}
		})
	}
}
