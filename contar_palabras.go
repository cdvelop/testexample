package main

import "strings"

// ContarPalabras retorna cuántas palabras hay en el string s.
// Una palabra se define como una secuencia de caracteres separados por espacios.
// Ejemplo: ContarPalabras("hola mundo go") == 3
func ContarPalabras(s string) int {
	return len(strings.Fields(s))
}
