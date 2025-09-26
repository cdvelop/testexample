package testexample

import (
	"fmt"
	"sort"
)

// SortInts retorna una copia del slice nums ordenado en forma ascendente.
// Ejemplo: SortInts([]int{3,1,2}) == []int{1,2,3}
func SortInts(nums []int) []int {
	// 1. Crear una copia del slice original.
	// Esto es crucial para no modificar el slice 'nums' pasado como argumento.
	// Se usa make para crear un nuevo slice con la misma longitud que 'nums'.
	copia := make([]int, len(nums))

	// Se usa copy para transferir los elementos de 'nums' a 'copia'.
	copy(copia, nums)

	// 2. Ordenar la copia.
	// El paquete 'sort' de Go proporciona la función Sort() para ordenar slices de enteros.
	sort.Ints(copia)

	// 3. Retornar el slice copiado y ordenado.
	return copia
}

func main() {
	fmt.Println("--- ✅ Caso 013 – Ordenar números (Pruebas) ---")

	// Slice de prueba original
	original := []int{10, 3, 45, 1, 8}

	// 1. Llamada a la función SortInts
	ordenado := SortInts(original)

	// 2. Muestra los resultados
	fmt.Printf("Slice Original: %v\n", original)
	fmt.Printf("Slice Ordenado: %v\n", ordenado)

	// Prueba de ejemplo del enunciado
	ejemplo := []int{3, 1, 2}
	ordenadoEjemplo := SortInts(ejemplo)
	fmt.Printf("\nEjemplo: SortInts(%v) == %v\n", ejemplo, ordenadoEjemplo)

	// Prueba con un slice ya ordenado
	yaOrdenado := []int{5, 6, 7}
	ordenadoYa := SortInts(yaOrdenado)
	fmt.Printf("\nYa Ordenado: SortInts(%v) == %v\n", yaOrdenado, ordenadoYa)
}
