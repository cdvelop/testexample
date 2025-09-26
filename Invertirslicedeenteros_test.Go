package test example

import "fmt"

// ReverseSlice retorna un nuevo slice con los elementos de nums en orden inverso.
// Ejemplo: ReverseSlice([]int{1,2,3}) == []int{3,2,1}
func ReverseSlice(nums []int) []int {
	// 1. Crear el nuevo slice de destino.
	// Usamos make para pre-asignar la memoria con la misma longitud que el slice original.
	// Esto es más eficiente que usar 'append' repetidamente.
	reversed := make([]int, len(nums))

	// 2. Iterar sobre el slice original.
	// Usaremos 'i' para el índice del slice original (de 0 a N-1)
	// y 'j' para el índice del nuevo slice (de N-1 a 0).
	// La relación es: reversed[j] = nums[i]
	
	for i := 0; i < len(nums); i++ {
		// Calcular el índice de destino 'j'
		// Si len(nums) es 3:
		// Cuando i=0, j = 3 - 1 - 0 = 2
		// Cuando i=1, j = 3 - 1 - 1 = 1
		// Cuando i=2, j = 3 - 1 - 2 = 0
		j := len(nums) - 1 - i
		
		// Asignar el elemento: el elemento en i se coloca en j.
		reversed[j] = nums[i]
	}

	// 3. Retornar el nuevo slice invertido.
	return reversed
}

func
() {
	fmt.Println("--- ✅ Caso 020 – Invertir Slice de Enteros (Pruebas) ---")

	// 1. Prueba de ejemplo del enunciado
	slice1 := []int{1, 2, 3}
	result1 := ReverseSlice(slice1)

	// 2. Prueba con un slice más largo
	slice2 := []int{10, 20, 30, 40, 50}
	result2 := ReverseSlice(slice2)

	// 3. Prueba con slice vacío
	slice3 := []int{}
	result3 := ReverseSlice(slice3)
    
	// 4. Prueba para verificar que el original no cambia
	original := []int{9, 8, 7}
	invertido := ReverseSlice(original)

	// Impresión de resultados
	fmt.Printf("ReverseSlice(%v) == %v (Esperado: [3 2 1])\n", slice1, result1)
	fmt.Printf("ReverseSlice(%v) == %v (Esperado: [50 40 30 20 10])\n", slice2, result2)
	fmt.Printf("ReverseSlice(%v) == %v (Esperado: [])\n", slice3, result3)
    fmt.Printf("\nOriginal: %v, Invertido: %v\n", original, invertido)
    fmt.Println("El slice original permanece inalterado, cumpliendo con el requisito de retornar un nuevo slice.")
}