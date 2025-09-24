package testexample

// Average retorna el promedio de una lista de enteros.
// Si la lista está vacía, debe retornar 0.
func Average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}
