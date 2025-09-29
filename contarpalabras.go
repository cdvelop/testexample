package certamen

import "strings"

// WordCount retorna cuántas palabras hay en el string s.
// Una palabra se define como una secuencia de caracteres separados por espacios.
// Ejemplo: WordCount("hola mundo go") == 3
func WordCount(s string) int {
	// strings.Fields separa la cadena por espacios, tabs o saltos de línea,
	// y retorna un slice con todas las "palabras".
	words := strings.Fields(s)
	return len(words)
}
