package testexample

import (
    "fmt"
    "math"
)

// Descripción:
// Un número primo es aquel mayor que 1 que solo tiene dos divisores: 1 y él mismo.
// Es decir, no puede dividirse exactamente (sin residuo) por ningún otro número.
//
// Ejemplos de números primos: 2, 3, 5, 7, 11, 13...
// Ejemplos de no primos: 1 (no cuenta como primo), 4 (2*2), 6 (2*3), 8 (2*4).
//
// IsPrime retorna true si n es un número primo.
// Ejemplo: IsPrime(7) == true
// Ejemplo: IsPrime(10) == false
func IsPrime(n int) bool {
    // 1. Manejo de casos triviales y no primos por definición.
    if n <= 1 {
        // Los números menores o iguales a 1 no son primos.
        return false
    }

    // 2. Manejo del 2 (el único número primo par).
    if n == 2 {
        return true
    }

    // 3. Excluir el resto de los números pares.
    if n%2 == 0 {
        // Cualquier número par mayor que 2 no es primo.
        return false
    }

    // 4. Bucle de optimización: Solo necesitamos verificar divisores impares hasta la raíz cuadrada de n.
    // Si un número tiene un divisor mayor que su raíz cuadrada, automáticamente tiene uno menor.
    // Usamos int(math.Sqrt(float64(n))) para obtener la raíz cuadrada.
    limit := int(math.Sqrt(float64(n)))
    for i := 3; i <= limit; i += 2 {
        if n%i == 0 {
            // Si encontramos algún divisor, el número no es primo.
            return false
        }
    }

    // 5. Si no se encontró ningún divisor, el número es primo.
    return true
}

🧪 Pruebas de Validación (Función main)
Para asegurarte de que tu función trabaja correctamente, aquí tienes la función main para probar varios casos, incluidos los ejemplos que se esperan.

Go

func main() {
    fmt.Println("--- ✅ Caso 012 – Número Primo (Pruebas) ---")

    // Pruebas de Números Primos
    fmt.Println("\n--- Primos (Expected: true) ---")
    primos := []int{2, 3, 5, 7, 13, 17, 97, 101}
    for _, num := range primos {
        // En un formato claro: IsPrime(7) -> true
        fmt.Printf("IsPrime(%d) -> %t\n", num, IsPrime(num))
    }

    // Pruebas de Números No Primos
    fmt.Println("\n--- No Primos (Expected: false) ---")
    noPrimos := []int{1, 4, 6, 9, 10, 15, 100, 10000}
    for _, num := range noPrimos {
        // En un formato claro: IsPrime(10) -> false
        fmt.Printf("IsPrime(%d) -> %t\n", num, IsPrime(num))
    }

    // Caso Límite
    fmt.Println("\n--- Caso Límite (Expected: false) ---")
    fmt.Printf("IsPrime(1) -> %t\n", 1, IsPrime(1)) // 1 no es primo por definición

    // Caso de No Primo grande que requiere la optimización de la raíz cuadrada
    fmt.Println("\n--- Prueba de Optimización (Expected: false) ---")
    numGrande := 121 // 11 * 11
    fmt.Printf("IsPrime(%d) -> %t (Raíz cuadrada de 121 es 11)\n", numGrande, IsPrime(numGrande)) 
}