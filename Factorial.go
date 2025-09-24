package testexample

import "fmt"

// Factorial calcula el factorial de un número entero no negativo n.
// Por convención: 0! = 1
func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// Pruebas de ejemplo
	fmt.Println(Factorial(0))  // 1
	fmt.Println(Factorial(5))  // 120
	fmt.Println(Factorial(6))  // 720
	fmt.Println(Factorial(10)) // 3628800
}
