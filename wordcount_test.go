package wordcount

import "testing"

func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "string vacío",
			input:    "",
			expected: 0,
		},
		{
			name:     "solo espacios",
			input:    "   ",
			expected: 0,
		},
		{
			name:     "una palabra",
			input:    "hola",
			expected: 1,
		},
		{
			name:     "dos palabras",
			input:    "hola mundo",
			expected: 2,
		},
		{
			name:     "tres palabras - caso ejemplo",
			input:    "hola mundo go",
			expected: 3,
		},
		{
			name:     "múltiples espacios entre palabras",
			input:    "hola    mundo    go",
			expected: 3,
		},
		{
			name:     "espacios al inicio",
			input:    "   hola mundo",
			expected: 2,
		},
		{
			name:     "espacios al final",
			input:    "hola mundo   ",
			expected: 2,
		},
		{
			name:     "espacios al inicio y final",
			input:    "   hola mundo go   ",
			expected: 3,
		},
		{
			name:     "una sola letra",
			input:    "a",
			expected: 1,
		},
		{
			name:     "muchas palabras",
			input:    "esto es una prueba con muchas palabras",
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordCount(tt.input)
			if result != tt.expected {
				t.Errorf("WordCount(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
