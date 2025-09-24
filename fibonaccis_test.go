package testexample

import "testing"

func TestFibonacci(t *testing.T) {
	casos := []struct {
		nombre string
		input  int
		espera int
	}{
		{"Caso base: F(0)", 0, 0},
		{"Caso base: F(1)", 1, 1},
		{"Caso chico: F(2)", 2, 1},
		{"Caso chico: F(3)", 3, 2},
		{"Caso chico: F(4)", 4, 3},
		{"Ejemplo clásico: F(5)", 5, 5},
		{"Ejemplo clásico: F(6)", 6, 8},
		{"Ejemplo clásico: F(7)", 7, 13},
		{"Medio: F(10)", 10, 55},
		{"Medio: F(12)", 12, 144},
		{"Más alto: F(15)", 15, 610},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			resultado := Fibonacci(c.input)
			if resultado != c.espera {
				t.Errorf("Fibonacci(%d) = %d; se esperaba %d",
					c.input, resultado, c.espera)
			}
		})
	}
}
