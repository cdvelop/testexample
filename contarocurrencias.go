package main // Paquete principal

// CountOccurrences retorna cuántas veces aparece target en el slice nums.
func CountOccurrences(nums []int, target int) int { // Cuenta ocurrencias de un número
	count := 0                 // Inicializa el contador
	for _, num := range nums { // Recorre cada elemento
		if num == target { // Si es igual al buscado
			count++ // Suma uno al contador
		}
	}
	return count // Retorna el total de ocurrencias
}
