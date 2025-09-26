package main

import "sort"

// SortInts retorna una copia del slice nums ordenado en forma ascendente.
func SortInts(nums []int) []int {
	// Creamos una copia del slice de entrada para no modificar el original.
	sliceCopia := make([]int, len(nums))
	copy(sliceCopia, nums)

	// Utilizamos la función Ints del paquete "sort" para ordenar la copia.
	sort.Ints(sliceCopia)

	return sliceCopia
}