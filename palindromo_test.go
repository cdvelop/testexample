package testexample

import (
	"testing"
)

// Helper function to reverse an integer and check if it's a palindrome
func isPalindrome(n int) bool {
	rev := 0
	cpy := n
	for n > 0 {
		rev = rev*10 + n%10
		n = n / 10
	}
	return rev == cpy
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{12321, true},
		{123, false},
		{1, true},
		{22, true},
		{10, false},
		{0, true},
	}

	for _, tt := range tests {
		result := isPalindrome(tt.input)
		if result != tt.expected {
			t.Errorf("isPalindrome(%d) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
