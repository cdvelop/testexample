package testexample

import (
	
	"strings"
)

// IsPalindrome retorna true si la palabra s es un palíndromo.
// Se ignoran mayúsculas para que "Ana" también sea válido.
func IsPalindrome(s string) bool {
	// Normalizamos a minúsculas
	s = strings.ToLower(s)

	// Usamos dos punteros, uno desde el inicio y otro desde el final
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}