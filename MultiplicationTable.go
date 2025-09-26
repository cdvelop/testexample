package main

// MultiplicationTable retorna un slice con los resultados de la tabla del número n, desde 1 hasta 10.
func MultiplicationTable(n int) []int {
	table := make([]int, 10)
	for i := 1; i <= 10; i++ {
		table[i-1] = n * i
	}
	return table
}
