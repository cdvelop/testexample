package testexample

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	// 2. Iterar a través de la mitad de la cadena
	for i := 0; i < len(s)/2; i++ {
		// 3. Comparar el carácter actual con el carácter opuesto
		if s[i] != s[len(s)-1-i] {
			return false // Si no son iguales, no es un palíndromo
		}
	}
	return true // Si el bucle termina, es un palíndromo
}
