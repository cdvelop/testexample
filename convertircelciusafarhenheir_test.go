package main

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name      string
		c         float64
		expectativa float64
	}{
		// Casos de prueba para la función CelsiusToFahrenheit.
		{"Caso 1: 0 Celsius es 32 Fahrenheit", 0, 32},
		{"Caso 2: 10 Celsius", 10, 50},
		{"Caso 3: 20 Celsius", 20, 68},
		{"Caso 4: 100 Celsius (punto de ebullición)", 100, 212},
		{"Caso 5: -40 Celsius (-40 Fahrenheit)", -40, -40},
		{"Caso 6: 15 Celsius", 15, 59},
		{"Caso 7: 30 Celsius", 30, 86},
		{"Caso 8: 37 Celsius (temperatura corporal)", 37, 98.6},
		{"Caso 9: -10 Celsius", -10, 14},
		{"Caso 10: 5 Celsius", 5, 41},
		{"Caso 11: 25 Celsius", 25, 77},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tt.c)
			if got != tt.expectativa {
				t.Errorf("CelsiusToFahrenheit(%f) = %f; se esperaba %f", tt.c, got, tt.expectativa)
			}
		})
	}
}