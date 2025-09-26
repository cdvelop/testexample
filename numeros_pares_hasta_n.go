package main

// NumerosParesHastaN retorna un slice con todos los números pares desde 0 hasta n (incluido).
func NumerosParesHastaN(n int) []int {
	if n < 0 {
		return []int{}
	}

	var evens []int
	for i := 0; i <= n; i += 2 {
		evens = append(evens, i)
	}
	return evens
}
