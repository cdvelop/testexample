package main

import "strings"

// LongestWord retorna la palabra más larga en el string s.
// En caso de empate, retorna la primera encontrada.
func LongestWord(s string) string {
	palabras := strings.Fields(s)
	if len(palabras) == 0 {
		return ""
	}

	palabraMasLarga := ""
	longitudMax := 0

	for _, palabra := range palabras {
		if len(palabra) > longitudMax {
			longitudMax = len(palabra)
			palabraMasLarga = palabra
		}
	}
	return palabraMasLarga
}