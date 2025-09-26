package wordcount

import "strings"

func WordCount(s string) int {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return 0
	}
	words := strings.Fields(trimmed)
	return len(words)
}
