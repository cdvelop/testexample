package wordcount

import (
	"reflect"
	"testing"
)

func TestMaxinSlice(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "ejemplo_del_enunciado",
			n:        6,
			expected: []int{0, 2, 4, 6},
		},
		{
			name:     "numero_cero",
			n:        0,
			expected: []int{0},
		},
		{
			name:     "numero_uno",
			n:        1,
			expected: []int{0},
		},
		{
			name:     "numero_dos",
			n:        2,
			expected: []int{0, 2},
		},
		{
			name:     "numero_impar_pequeno",
			n:        3,
			expected: []int{0, 2},
		},
		{
			name:     "numero_par_pequeno",
			n:        4,
			expected: []int{0, 2, 4},
		},
		{
			name:     "numero_impar_mediano",
			n:        9,
			expected: []int{0, 2, 4, 6, 8},
		},
		{
			name:     "numero_par_mediano",
			n:        10,
			expected: []int{0, 2, 4, 6, 8, 10},
		},
		{
			name:     "numero_grande_par",
			n:        20,
			expected: []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		},
		{
			name:     "numero_grande_impar",
			n:        15,
			expected: []int{0, 2, 4, 6, 8, 10, 12, 14},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvenNumbersUntil(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("EvenNumbersUntil(%d) = %v, esperado %v",
					tt.n, result, tt.expected)
			}
		})
	}
}
