package testexample

import "fmt"

// EvenNumbersUntil retorna un slice con todos los números pares desde 0 hasta n (incluido).
// Ejemplo: EvenNumbersUntil(6) == []int{0,2,4,6}
func EvenNumbersUntil(n int) []int {
	// Inicializamos un slice vacío para almacenar los números pares.
	pares := []int{}

	// Bucle: Comenzamos en 0 y continuamos mientras el número sea menor o igual a n.
	// La clave es el incremento: i += 2 asegura que solo trabajemos con números pares.
	for i := 0; i <= n; i += 2 {
		// Añadimos el número par actual al slice.
		pares = append(pares, i)
	}

	// Retornamos el slice resultante.
	return pares
}

func () {
	fmt.Println("--- ✅ Caso 016 – Números pares hasta N (Pruebas) ---")

	// 1. Prueba de ejemplo del enunciado
	n1 := 6
	result1 := EvenNumbersUntil(n1)

	// 2. Prueba con un número impar grande (el slice debe terminar en n-1)
	n2 := 11
	result2 := EvenNumbersUntil(n2)

	// 3. Prueba con n=0
	n3 := 0
	result3 := EvenNumbersUntil(n3)

	// 4. Prueba con n=1 (solo incluye el 0)
	n4 := 1
	result4 := EvenNumbersUntil(n4)

	// Impresión de resultados
	fmt.Printf("EvenNumbersUntil(%d) == %v (Esperado: [0 2 4 6])\n", n1, result1)
	fmt.Printf("EvenNumbersUntil(%d) == %v (Esperado: [0 2 4 6 8 10])\n", n2, result2)
	fmt.Printf("EvenNumbersUntil(%d) == %v (Esperado: [0])\n", n3, result3)
	fmt.Printf("EvenNumbersUntil(%d) == %v (Esperado: [0])\n", n4, result4)
}
