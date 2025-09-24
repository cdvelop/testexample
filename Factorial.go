package testexample
// Factorial calcula el factorial de un número entero no negativo.
// La función se implementa en forma recursiva.
// Ejemplo: Factorial(5) = 120 (5*4*3*2*1).
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}