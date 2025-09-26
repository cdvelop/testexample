package main // Paquete principal

import "sort" // Importa el paquete sort para ordenar slices

// SortInts retorna una copia del slice nums ordenado en forma ascendente.
func SortInts(nums []int) []int { // Función que ordena un slice de enteros
	sliceCopia := make([]int, len(nums)) // Crea un nuevo slice con la misma longitud
	copy(sliceCopia, nums)               // Copia los elementos del slice original
	sort.Ints(sliceCopia)                // Ordena el slice copia de menor a mayor
	return sliceCopia                    // Retorna el slice ordenado
}
