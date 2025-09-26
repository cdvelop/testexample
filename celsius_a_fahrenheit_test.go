package main

import (
	"math"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name        string
		celsius     float64
		expectativa float64
	}{
		{"Caso 1: Punto de congelación", 0, 32},
		{"Caso 2: Punto de ebullición", 100, 212},
		{"Caso 3: Temperatura corporal", 37, 98.6},
		{"Caso 4: Temperatura negativa", -10, 14},
		{"Caso 5: Punto donde C y F son iguales", -40, -40},
		{"Caso 6: Día caluroso", 30, 86},
		{"Caso 7: Día fresco", 15, 59},
		{"Caso 8: Con decimales", 25.5, 77.9},
		{"Caso 9: Cero absoluto (aprox)", -273.15, -459.67},
		{"Caso 10: Otro valor negativo", -20, -4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tt.celsius)
			// Comparamos con una pequeña tolerancia para evitar errores de precisión en floats
			if math.Abs(got-tt.expectativa) > 0.001 {
				t.Errorf("CelsiusToFahrenheit(%.2f) = %.2f; want %.2f", tt.celsius, got, tt.expectativa)
			}
		})
	}
}
