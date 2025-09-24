package testexample

import "regexp"

func isPalindrome(s string) bool {
	pattern := regexp.MustCompile(`^[\p{L}]+$`)
	if !pattern.MatchString(s) {
		return false
	}
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}
