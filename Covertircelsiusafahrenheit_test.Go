package testexample

import "fmt"

// CelsiusToFahrenheit convierte grados Celsius a Fahrenheit.
// Fórmula: F = C * 9/5 + 32
// Ejemplo: CelsiusToFahrenheit(0) == 32
func CelsiusToFahrenheit(c float64) float64 {
	// Aplicamos la fórmula: F = C * (9.0/5.0) + 32.0
	// Usamos 9.0/5.0 para garantizar que la división se realice con precisión flotante
	// y 32.0 para mantener la coherencia del tipo float64 en toda la operación.
	fahrenheit := c*(9.0/5.0) + 32.0

	return fahrenheit
}

func () {
	fmt.Println("--- ✅ Caso 015 – Convertir Celsius a Fahrenheit (Pruebas) ---")

	// 1. Pruebas con valores conocidos

	// Punto de congelación del agua: 0°C = 32°F
	c1 := 0.0
	f1 := CelsiusToFahrenheit(c1)

	// Temperatura ambiente aproximada: 20°C = 68°F
	c2 := 20.0
	f2 := CelsiusToFahrenheit(c2)

	// Punto de ebullición del agua: 100°C = 212°F
	c3 := 100.0
	f3 := CelsiusToFahrenheit(c3)

	// Temperatura negativa: -40°C = -40°F (El punto donde coinciden)
	c4 := -40.0
	f4 := CelsiusToFahrenheit(c4)

	// 2. Impresión de resultados
	fmt.Printf("%.1f°C es igual a %.1f°F (Esperado: 32.0)\n", c1, f1)
	fmt.Printf("%.1f°C es igual a %.1f°F (Esperado: 68.0)\n", c2, f2)
	fmt.Printf("%.1f°C es igual a %.1f°F (Esperado: 212.0)\n", c3, f3)
	fmt.Printf("%.1f°C es igual a %.1f°F (Esperado: -40.0)\n", c4, f4)
}
