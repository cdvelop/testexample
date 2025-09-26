package main

import "strings"

// LongestWord retorna la palabra más larga en el string s.
// En caso de empate, retorna la primera encontrada.
func LongestWord(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}

	longest := words[0]
	for i := 1; i < len(words); i++ {
		if len(words[i]) > len(longest) {
			longest = words[i]
		}
	}
	return longest
}
