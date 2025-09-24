package testexample

// CountVowels retorna cuántas vocales (a,e,i,o,u) hay en el string.
func ContarVocales(s string) int {
	count := 0
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			count++
		}
	}
	return count
}
