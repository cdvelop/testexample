package testexample

import "testing"

func TestIsPalindrome(t *testing.T) {
	casos := []struct {
		nombre string
		input  string
		espera bool
	}{
		{"Cadena vacía", "", true},
		{"Un solo carácter", "a", true},
		{"Dos letras iguales", "aa", true},
		{"Dos letras distintas", "ab", false},
		{"Palíndromo simple", "radar", true},
		{"Palíndromo par", "abba", true},
		{"No palíndromo corto", "hola", false},
		{"Palíndromo con mayúsculas", "Ana", true},
		{"No palíndromo con mayúsculas", "GoLang", false},
		{"Palíndromo largo", "reconocer", true},
		{"No palíndromo largo", "universidad", false},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			resultado := IsPalindrome(c.input)
			if resultado != c.espera {
				t.Errorf("IsPalindrome(%q) = %v; esperado %v",
					c.input, resultado, c.espera)
			}
		})
	}
}
