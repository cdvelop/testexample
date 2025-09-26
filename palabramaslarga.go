package main // Paquete principal

import "strings" // Importa el paquete strings para manipular cadenas

// LongestWord retorna la palabra más larga en el string s.
// En caso de empate, retorna la primera encontrada.
func LongestWord(s string) string { // Busca la palabra más larga
	palabras := strings.Fields(s) // Separa el string en palabras
	if len(palabras) == 0 {       // Si no hay palabras, retorna vacío
		return ""
	}
	palabraMasLarga := ""              // Guarda la palabra más larga
	longitudMax := 0                   // Guarda la longitud máxima
	for _, palabra := range palabras { // Recorre cada palabra
		if len(palabra) > longitudMax { // Si es más larga
			longitudMax = len(palabra) // Actualiza longitud máxima
			palabraMasLarga = palabra  // Actualiza palabra más larga
		}
	}
	return palabraMasLarga // Retorna la palabra más larga
}
