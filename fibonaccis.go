package testexample


// Fibonacci retorna el número de Fibonacci en la posición n.
// Se usa un enfoque iterativo porque es más eficiente que la recursión
// cuando n crece un poco más.
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}