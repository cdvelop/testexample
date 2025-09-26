package wordcount

import "testing"

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "ejemplo_basico_del_enunciado",
			nums:     []int{1, 2, 2, 3, 2},
			target:   2,
			expected: 3,
		},
		{
			name:     "slice_vacio",
			nums:     []int{},
			target:   5,
			expected: 0,
		},
		{
			name:     "elemento_no_existe",
			nums:     []int{1, 3, 5, 7, 9},
			target:   2,
			expected: 0,
		},
		{
			name:     "todos_elementos_son_target",
			nums:     []int{7, 7, 7, 7},
			target:   7,
			expected: 4,
		},
		{
			name:     "un_solo_elemento_que_coincide",
			nums:     []int{42},
			target:   42,
			expected: 1,
		},
		{
			name:     "un_solo_elemento_que_no_coincide",
			nums:     []int{10},
			target:   20,
			expected: 0,
		},
		{
			name:     "target_al_inicio_y_final",
			nums:     []int{5, 1, 2, 3, 5},
			target:   5,
			expected: 2,
		},
		{
			name:     "numeros_negativos_con_target_negativo",
			nums:     []int{-1, -2, -1, 0, -1},
			target:   -1,
			expected: 3,
		},
		{
			name:     "numeros_negativos_target_positivo",
			nums:     []int{-5, -3, 2, -1, 2, 2},
			target:   2,
			expected: 3,
		},
		{
			name:     "slice_grande_con_muchas_ocurrencias",
			nums:     []int{1, 9, 9, 2, 9, 3, 9, 4, 9, 9, 5, 9},
			target:   9,
			expected: 7,
		},
		{
			name:     "target_cero_en_slice_con_ceros",
			nums:     []int{0, 1, 0, 2, 0, 0, 3},
			target:   0,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountOccurrences(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("CountOccurrences(%v, %d) = %d, esperado %d",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
