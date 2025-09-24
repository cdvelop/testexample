package testexample

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"radar", true},
		{"civic", true},
		{"level", true},
		{"racecar", true},
		{"noon", true},
		{"hello", false},
		{"world", false},
		{"madam", true},
		{"python", false},
		{"java", false},
	}

	for _, tc := range tests {
		result := isPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("Para '%s', se esperaba %t pero se obtuvo %t", tc.input, tc.expected, result)
		}
	}
}
