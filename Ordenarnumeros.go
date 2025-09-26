package testexample

import (
	"fmt"
	"sort"
)

// SortInts retorna una copia del slice nums ordenado en forma ascendente.
// Ejemplo: SortInts([]int{3,1,2}) == []int{1,2,3}
func SortInts(nums []int) []int {
	// Crear una copia del slice original
	copia := make([]int, len(nums))
	copy(copia, nums)

	// Ordenar la copia en forma ascendente
	sort.Ints(copia)

	return copia
}

func main() {
	// Ejemplo de uso
	nums := []int{3, 1, 2}
	fmt.Println("Original:", nums)
	fmt.Println("Ordenado:", SortInts(nums))
}
