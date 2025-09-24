package testexample

import "testing"

func TestFactorial(t *testing.T) {
	casos := []struct {
		nombre string
		input  int
		espera int
	}{
		{"Caso base: 0!", 0, 1},
		{"Caso base: 1!", 1, 1},
		{"Valor pequeño: 2!", 2, 2},
		{"Valor pequeño: 3!", 3, 6},
		{"Medio: 4!", 4, 24},
		{"Medio: 5!", 5, 120},
		{"Clásico: 6!", 6, 720},
		{"Medio-alto: 7!", 7, 5040},
		{"Medio-alto: 8!", 8, 40320},
		{"Alto: 10!", 10, 3628800},
		{"Mayor: 12!", 12, 479001600},
		{"Mayor: 15!", 15, 1307674368000},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			resultado := Factorial(c.input)
			if resultado != c.espera {
				t.Errorf("Factorial(%d) = %d; se esperaba %d",
					c.input, resultado, c.espera)
			}
		})
	}
}
