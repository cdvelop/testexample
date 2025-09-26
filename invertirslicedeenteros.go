package main

// ReverseSlice retorna un nuevo slice con los elementos de nums en orden inverso.
func ReverseSlice(nums []int) []int {
	// Creamos un nuevo slice con la misma longitud que el original.
	reversed := make([]int, len(nums))

	// Recorremos el slice original y copiamos los elementos en el nuevo
	// slice en orden inverso.
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		reversed[i] = nums[j]
	}

	return reversed
}