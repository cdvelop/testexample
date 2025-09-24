package testexample

import "testing"

func TestSumPositives(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Solo positivos", []int{1, 2, 3}, 6},
		{"Solo negativos", []int{-1, -2, -3}, 0},
		{"Mixto positivos y negativos", []int{-1, 2, 3, 0}, 5},
		{"Incluye cero", []int{0, 0, 0}, 0},
		{"Un solo positivo", []int{5}, 5},
		{"Un solo negativo", []int{-5}, 0},
		{"Slice vacío", []int{}, 0},
		{"Grandes positivos", []int{100, 200, 300}, 600},
		{"Positivos y ceros", []int{0, 1, 0, 2}, 3},
		{"Negativos y ceros", []int{-1, 0, -2, 0}, 0},
		{"Todos iguales positivos", []int{7, 7, 7}, 21},
		{"Todos iguales negativos", []int{-7, -7, -7}, 0},
		{"Positivos repetidos", []int{2, 2, 2, -2}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumPositives(tt.input)
			if got != tt.expected {
				t.Errorf("SumPositives(%v) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}
