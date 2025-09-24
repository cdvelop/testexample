package testexample

// SumPositives retorna la suma de los enteros positivos dentro del slice nums.
// Ejemplo: SumPositives([]int{-1, 2, 3, 0}) == 5
// Si no hay positivos, debe retornar 0.
func SumPositives(nums []int) int {
	sum := 0
	for _, n := range nums {
		if n > 0 {
			sum += n
		}
	}
	return sum
}
