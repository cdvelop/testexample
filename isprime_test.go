package wordcount

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "cero no es primo",
			input:    0,
			expected: false,
		},
		{
			name:     "uno no es primo",
			input:    1,
			expected: false,
		},
		{
			name:     "dos es primo",
			input:    2,
			expected: true,
		},
		{
			name:     "tres es primo",
			input:    3,
			expected: true,
		},
		{
			name:     "cuatro no es primo",
			input:    4,
			expected: false,
		},
		{
			name:     "cinco es primo",
			input:    5,
			expected: true,
		},
		{
			name:     "seis no es primo",
			input:    6,
			expected: false,
		},
		{
			name:     "siete es primo - caso ejemplo",
			input:    7,
			expected: true,
		},
		{
			name:     "nueve no es primo",
			input:    9,
			expected: false,
		},
		{
			name:     "diez no es primo - caso ejemplo",
			input:    10,
			expected: false,
		},
		{
			name:     "once es primo",
			input:    11,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.input)
			if result != tt.expected {
				t.Errorf("IsPrime(%d) = %t, expected %t", tt.input, result, tt.expected)
			}
		})
	}
}
