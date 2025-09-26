package main // Paquete principal

import "strings" // Importa el paquete strings para manipular cadenas

// WordCount retorna cuántas palabras hay en el string s.
func WordCount(s string) int { // Cuenta palabras en un string
	return len(strings.Fields(s)) // Separa por espacios y cuenta
}
