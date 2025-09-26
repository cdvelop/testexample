package main // Paquete principal

// ReverseSlice retorna un nuevo slice con los elementos de nums en orden inverso.
func ReverseSlice(nums []int) []int { // Invierte un slice de enteros
	reversed := make([]int, len(nums))                           // Crea un nuevo slice con la misma longitud
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 { // Recorre desde ambos extremos
		reversed[i] = nums[j] // Asigna el valor inverso
	}
	return reversed // Retorna el slice invertido
}
