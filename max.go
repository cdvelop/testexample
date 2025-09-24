package main

// Max retorna el mayor de los dos números a y b.
// Ejemplo: Max(3, 7) == 7
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}