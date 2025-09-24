package testexample

// Max retorna el mayor de los dos números a y b.
// Ejemplo: Max(3, 7) == 7
func Max(a, b int) int {
	// TODO: implementar
	return 0
}
func TestMax_BasicCases(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"a_less_than_b", 3, 7, 7},
		{"a_greater_than_b", 10, 2, 10},
		{"a_equal_b", 5, 5, 5},
		{"negative_and_positive", -4, 2, 2},
		{"both_negative_a_less", -10, -3, -3},
		{"both_negative_a_greater", -2, -8, -2},
		{"zero_and_positive", 0, 9, 9},
		{"zero_and_negative", 0, -1, 0},
		{"large_values_a", 1 << 30, 1 << 29, 1 << 30},
		{"large_values_b", 1 << 29, 1 << 30, 1 << 30},
	}

	for _, tc := range tests {
		tc := tc // captura local para subtests
		t.Run(tc.name, func(t *testing.T) {
			got := Max(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("Max(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestMax_CommutativeProperty(t *testing.T) {
	// Max should be commutativa: Max(a,b) == Max(b,a)
	cases := []struct {
		a int
		b int
	}{
		{3, 7},
		{-5, 2},
		{0, 0},
		{-10, -10},
		{123456, 654321},
	}

	for i, c := range cases {
		t.Run("case#"+string(rune(i+'A')), func(t *testing.T) {
			if Max(c.a, c.b) != Max(c.b, c.a) {
				t.Fatalf("Commutative property failed for (%d,%d): %d vs %d", c.a, c.b, Max(c.a, c.b), Max(c.b, c.a))
			}
		})
	}
}

// Example function to demonstrate expected behavior in Go doc tests.
// It is optional, but ayuda a mostrar el uso en la documentación.
func ExampleMax() {
	_ = Max(3, 7) // expected 7 (no output in example test)
	// To see Example tests pass, run: go test
}
