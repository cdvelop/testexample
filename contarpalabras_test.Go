package testexample

import (
	"fmt"
	"strings"
)

// WordCount retorna cuántas palabras hay en el string s.
// Una palabra se define como una secuencia de caracteres separados por espacios.
func WordCount(s string) int {
	// La función strings.Fields() es la más robusta para este caso.
	// Divide el string en un slice de sub-cadenas,
	// separando por uno o más espacios en blanco, tabulaciones o saltos de línea.
	palabras := strings.Fields(s)

	// La función len() retorna la cantidad de elementos en el slice.
	return len(palabras)
}

func contarpalabras() {
	// ## ✅ Caso 011 – Contar palabras (5 pts.)
	fmt.Println("--- Pruebas de WordCount ---")

	// Caso 1: Texto simple.
	texto1 := "Hola mundo, esto es una prueba."
	fmt.Printf("El texto '%s' tiene %d palabras.\n", texto1, WordCount(texto1))
	// Salida esperada: 5

	// Caso 2: Texto con espacios adicionales.
	texto2 := "  Hola   mundo,   con   muchos   espacios. "
	fmt.Printf("El texto '%s' tiene %d palabras.\n", texto2, WordCount(texto2))
	// Salida esperada: 5

	// Caso 3: Texto con puntuación y saltos de línea.
	texto3 := "Este es un texto,\nque tiene\nvarias líneas."
	fmt.Printf("El texto '%s' tiene %d palabras.\n", texto3, WordCount(texto3))
	// Salida esperada: 7

	// Caso 4: String vacío.
	texto4 := ""
	fmt.Printf("El texto '%s' tiene %d palabras.\n", texto4, WordCount(texto4))
	// Salida esperada: 0
}
