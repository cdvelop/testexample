package testexample

import "fmt"

// MultiplicationTable retorna un slice con los resultados de la tabla del número n, desde 1 hasta 10.
// Ejemplo: MultiplicationTable(3) == []int{3,6,9,12,15,18,21,24,27,30}
func MultiplicationTable(n int) []int {
	// 1. Crear un slice para almacenar los resultados.
	// Sabemos que contendrá exactamente 10 elementos, por lo que podemos inicializarlo con esa capacidad.
	tabla := make([]int, 0, 10)

	// 2. Iterar desde el 1 hasta el 10 (incluido).
	for i := 1; i <= 10; i++ {
		// 3. Calcular el resultado (n * i).
		resultado := n * i
		
		// 4. Añadir el resultado al slice.
		tabla = append(tabla, resultado)
	}

	// 5. Retornar el slice completo de la tabla.
	return tabla
}

func
() {
	fmt.Println("--- ✅ Caso 019 – Tabla de Multiplicar (Pruebas) ---")

	// 1. Prueba de ejemplo del enunciado (Tabla del 3)
	n1 := 3
	result1 := MultiplicationTable(n1)
	
	// 2. Prueba con la Tabla del 7
	n2 := 7
	result2 := MultiplicationTable(n2)

	// 3. Prueba con la Tabla del 1
	n3 := 1
	result3 := MultiplicationTable(n3)

	// Impresión de resultados
	fmt.Printf("MultiplicationTable(%d) == %v\n", n1, result1)
	fmt.Printf("MultiplicationTable(%d) == %v\n", n2, result2)
	fmt.Printf("MultiplicationTable(%d) == %v\n", n3, result3)
}