package wordcount

import "sort"

func SortInts(nums []int) []int {
	copia := make([]int, len(nums))
	copy(copia, nums)
	sort.Ints(copia)
	return copia
}
