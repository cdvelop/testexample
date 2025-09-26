package testexample

import "fmt"

// CountOccurrences retorna cuántas veces aparece target en el slice nums.
// Ejemplo: CountOccurrences([]int{1, 2, 3, 2, 4}, 2) == 2
func CountOccurrences(nums []int, target int) int {
	// Inicializamos un contador a cero. Este será nuestro resultado.
	count := 0

	// Iteramos sobre cada número (num) en el slice 'nums'.
	for _, num := range nums {
		// Comparamos el número actual con el valor objetivo.
		if num == target {
			// Si coinciden, incrementamos el contador.
			count++
		}
	}

	// Devolvemos el conteo final de ocurrencias.
	return count
}

func ai() {
	fmt.Println("--- 

	// 1. Slice de prueba
	data := []int{5, 10, 5, 15, 20, 5, 10, 5, 25}

	// 2. Pruebas con diferentes valores objetivo
	target1 := 5 // Ocurre 4 veces
	result1 := CountOccurrences(data, target1)

	target2 := 10 // Ocurre 2 veces
	result2 := CountOccurrences(data, target2)

	target3 := 30 // Ocurre 0 veces
	result3 := CountOccurrences(data, target3)

	// 3. Impresión de resultados
	fmt.Printf("Slice de datos: %v\n", data)
	fmt.Printf("\nTarget: %d, Ocurrencias: %d (Esperado: 4)\n", target1, result1)
	fmt.Printf("Target: %d, Ocurrencias: %d (Esperado: 2)\n", target2, result2)
	fmt.Printf("Target: %d, Ocurrencias: %d (Esperado: 0)\n", target3, result3)
}
