package main

import "sort"

// SortInts retorna una copia del slice nums ordenado en forma ascendente.
func SortInts(nums []int) []int {
	// Crear una copia para no modificar el slice original
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)

	sort.Ints(sortedNums)
	return sortedNums
}
