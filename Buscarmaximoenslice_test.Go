package testexample

import "fmt"

// MaxInSlice retorna el valor máximo dentro de nums.
// Si el slice está vacío, debe retornar 0.
// Ejemplo: MaxInSlice([]int{3,7,2}) == 7
func MaxInSlice(nums []int) int {
	// 1. Manejar el caso de slice vacío.
	if len(nums) == 0 {
		return 0
	}

	// 2. Inicializar el valor máximo.
	// Asumimos que el primer elemento es el máximo inicialmente.
	maxVal := nums[0]

	// 3. Iterar a través del resto del slice (empezando desde el segundo elemento).
	// Usamos el índice `i` para claridad, pero también podríamos usar el rango completo.
	for i := 1; i < len(nums); i++ {
		// 4. Comparar el valor actual con el máximo registrado.
		if nums[i] > maxVal {
			// Si el valor actual es mayor, actualizamos maxVal.
			maxVal = nums[i]
		}
	}

	// 5. Retornar el valor máximo encontrado.
	return maxVal
}

func 
() {
	fmt.Println("--- ✅ Caso 017 – Buscar Máximo en Slice (Pruebas) ---")

	// 1. Prueba de ejemplo del enunciado
	slice1 := []int{3, 7, 2}
	result1 := MaxInSlice(slice1)

	// 2. Prueba con números desordenados (máximo al principio)
	slice2 := []int{90, 15, 7, 30}
	result2 := MaxInSlice(slice2)

	// 3. Prueba con números negativos
	slice3 := []int{-10, -5, -20, -1}
	result3 := MaxInSlice(slice3)

	// 4. Prueba de caso límite: slice vacío
	slice4 := []int{}
	result4 := MaxInSlice(slice4)

	// 5. Prueba de caso con un solo elemento
	slice5 := []int{42}
	result5 := MaxInSlice(slice5)

	// Impresión de resultados
	fmt.Printf("MaxInSlice(%v) == %d (Esperado: 7)\n", slice1, result1)
	fmt.Printf("MaxInSlice(%v) == %d (Esperado: 90)\n", slice2, result2)
	fmt.Printf("MaxInSlice(%v) == %d (Esperado: -1)\n", slice3, result3)
	fmt.Printf("MaxInSlice(%v) == %d (Esperado: 0)\n", slice4, result4)
	fmt.Printf("MaxInSlice(%v) == %d (Esperado: 42)\n", slice5, result5)
}
