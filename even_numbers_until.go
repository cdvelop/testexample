package wordcount

func EvenNumbersUntil(n int) []int {
	var result []int
	for i := 0; i <= n; i += 2 {
		result = append(result, i)

	}
	return result
}
