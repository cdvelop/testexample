// Estudiante: [CatalinaToledo]
// Curso: Programación en Go
// Nota: Implementación sencilla y comentada, estilo estudiante.

package main

import "fmt"

// SumPositives retorna la suma de los enteros positivos dentro del slice nums.
// Ejemplo: SumPositives([]int{-1, 2, 3, 0}) == 5
// Si no hay positivos, debe retornar 0.
func SumPositives(nums []int) int {
	suma := 0
	for _, n := range nums {
		if n > 0 { // contamos sólo positivos estrictos
			suma += n
		}
	}
	return suma
}

func main() {
	// Ejemplos de uso (para probar manualmente)
	casos := [][]int{
		{-1, 2, 3, 0},
		{-5, -2, 0},
		{10, 20, 30},
		{},
		{0, 0, 0},
	}
	for _, c := range casos {
		fmt.Printf("SumPositives(%v) = %d\n", c, SumPositives(c))
	}
}
