package testexample

import "testing"

func TestAverage(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{[]int{2, 4, 6}, 4.0},         // 12/3 = 4
		{[]int{1, 2, 3, 4, 5}, 3.0},   // 15/5 = 3
		{[]int{}, 0.0},                // lista vacía
		{[]int{10}, 10.0},             // un solo número
		{[]int{-2, -4, -6}, -4.0},     // promedio negativo
		{[]int{1, -1, 1, -1}, 0.0},    // mezcla de positivos y negativos
		{[]int{0, 0, 0}, 0.0},         // todos ceros
		{[]int{5, 15, 25}, 15.0},      // 45/3 = 15
		{[]int{100, 200, 300}, 200.0}, // 600/3 = 200
		{[]int{7, 14, 21, 28}, 17.5},  // 70/4 = 17.5
	}

	for _, test := range tests {
		result := Average(test.input)
		if result != test.expected {
			t.Errorf("Average(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
