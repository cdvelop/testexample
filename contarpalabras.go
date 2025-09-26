package main

import "strings"

// WordCount retorna cuántas palabras hay en el string s.
func WordCount(s string) int {
	return len(strings.Fields(s))
}