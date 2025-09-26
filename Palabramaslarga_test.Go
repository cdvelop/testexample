package testexample

import (
	"fmt"
	"strings"
)

// LongestWord retorna la palabra más larga en el string s.
// En caso de empate, retorna la primera encontrada.
// Ejemplo: LongestWord("go es divertido") == "divertido"
func LongestWord(s string) string {
	// 1. Dividir el string en un slice de palabras.
	// strings.Fields() maneja de forma robusta los espacios en blanco múltiples.
	words := strings.Fields(s)

	// 2. Manejar el caso de string vacío o sin palabras.
	if len(words) == 0 {
		return ""
	}

	// 3. Inicializar la palabra más larga como la primera palabra.
	longestWord := words[0]

	// 4. Iterar sobre el resto de las palabras.
	// Comenzamos desde el segundo elemento (índice 1).
	for i := 1; i < len(words); i++ {
		currentWord := words[i]

		// 5. Comparar longitudes.
		// Si la palabra actual es más larga que la palabra más larga registrada,
		// la reemplazamos.
		// NOTA: Como el problema pide la *primera* en caso de empate,
		// usamos el operador de comparación estricto ">" (mayor que)
		// y no ">=" (mayor o igual que).
		if len(currentWord) > len(longestWord) {
			longestWord = currentWord
		}
	}

	// 6. Retornar la palabra más larga.
	return longestWord
}

func
() {
	fmt.Println("--- ✅ Caso 018 – Palabra Más Larga (Pruebas) ---")

	// 1. Prueba de ejemplo del enunciado
	s1 := "go es divertido"
	result1 := LongestWord(s1)

	// 2. Prueba con empate (debe retornar "palabra")
	s2 := "palabra larga palabra" 
	result2 := LongestWord(s2)

	// 3. Prueba con puntuación y múltiples espacios
	s3 := "  La programación, es genial!  "
	result3 := LongestWord(s3) // Retornará "programación," o "genial!"

	// 4. Prueba con string vacío
	s4 := ""
	result4 := LongestWord(s4)
    
	// 5. Prueba con una sola palabra
	s5 := "Unica"
	result5 := LongestWord(s5)

	// Impresión de resultados
	fmt.Printf("LongestWord(\"%s\") == \"%s\" (Esperado: \"divertido\")\n", s1, result1)
	fmt.Printf("LongestWord(\"%s\") == \"%s\" (Esperado: \"palabra\")\n", s2, result2)
	fmt.Printf("LongestWord(\"%s\") == \"%s\" (Esperado: \"programación,\")\n", s3, result3)
	fmt.Printf("LongestWord(\"%s\") == \"%s\" (Esperado: \"\")\n", s4, result4)
	fmt.Printf("LongestWord(\"%s\") == \"%s\" (Esperado: \"Unica\")\n", s5, result5)
}