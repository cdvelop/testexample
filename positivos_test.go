func TestSumPositives(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"Mixto", []int{-1, 2, 3, 0}, 5},
		{"Solo negativos y ceros", []int{-5, -1, 0}, 0},
		{"Solo positivos", []int{1, 2, 3}, 6},
		{"Vacío", []int{}, 0},
		{"Solo ceros", []int{0, 0, 0}, 0},
		{"Un positivo", []int{-2, 7, -3}, 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SumPositives(tc.input)
			if got != tc.want {
				t.Fatalf("SumPositives(%v) = %d; quiero %d", tc.input, got, tc.want)
			}
		})
	}
}
