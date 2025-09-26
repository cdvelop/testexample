package wordcount

import (
	"math"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{
			name:     "ejemplo_del_enunciado",
			celsius:  0,
			expected: 32,
		},
		{
			name:     "punto_de_ebullicion_del_agua",
			celsius:  100,
			expected: 212,
		},
		{
			name:     "temperatura_negativa",
			celsius:  -10,
			expected: 14,
		},
		{
			name:     "cero_absoluto",
			celsius:  -273.15,
			expected: -459.67,
		},
		{
			name:     "temperatura_corporal",
			celsius:  37,
			expected: 98.6,
		},
		{
			name:     "temperatura_ambiente",
			celsius:  25,
			expected: 77,
		},
		{
			name:     "numero_decimal_positivo",
			celsius:  36.5,
			expected: 97.7,
		},
		{
			name:     "numero_decimal_negativo",
			celsius:  -5.5,
			expected: 22.1,
		},
		{
			name:     "temperatura_muy_alta",
			celsius:  1000,
			expected: 1832,
		},
		{
			name:     "temperatura_muy_baja",
			celsius:  -50,
			expected: -58,
		},
		{
			name:     "punto_de_fusion_del_hielo",
			celsius:  0.0,
			expected: 32.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CelsiusToFahrenheit(tt.celsius)

			// Usamos una tolerancia de 0.01 para la comparación
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("CelsiusToFahrenheit(%.2f) = %.2f, esperado %.2f",
					tt.celsius, result, tt.expected)
			}
		})
	}
}
