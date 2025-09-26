package wordcount

import (
	"reflect"
	"testing"
)

func TestSortInts(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "slice vacío",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "slice con un solo elemento",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "slice ya ordenado",
			input:    []int{1, 3, 2, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "slice ordenado en reversa",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "slice desordenado",
			input:    []int{3, 1, 4, 1, 5},
			expected: []int{1, 1, 3, 4, 5},
		},
		{
			name:     "slice con números negativos",
			input:    []int{-5, -1, -10, 0},
			expected: []int{-10, -5, -1, 0},
		},
		{
			name:     "slice con números duplicados",
			input:    []int{7, 3, 7, 3, 1},
			expected: []int{1, 3, 3, 7, 7},
		},
		{
			name:     "slice con números aleatorios grandes",
			input:    []int{99, 10, 50, 2, 77, 100},
			expected: []int{2, 10, 50, 77, 99, 100},
		},
		{
			name:     "slice con números negativos y positivos",
			input:    []int{1, -2, 3, -4, 5, 0},
			expected: []int{-4, -2, 0, 1, 3, 5},
		},
		{
			name:     "slice con un gran número de elementos",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "otro slice desordenado",
			input:    []int{20, 10, 30, 5, 25},
			expected: []int{5, 10, 20, 25, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortInts(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SortInts(%v) = %v, se esperaba %v", tt.input, result, tt.expected)
			}
		})
	}
}
