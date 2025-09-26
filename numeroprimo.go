package main // Paquete principal

// IsPrime retorna true si n es un número primo.
func IsPrime(n int) bool { // Verifica si un número es primo
	if n <= 1 { // 0 y 1 no son primos
		return false
	}
	for i := 2; i*i <= n; i++ { // Prueba divisores hasta raíz de n
		if n%i == 0 { // Si es divisible, no es primo
			return false
		}
	}
	return true // Si no encontró divisores, es primo
}
