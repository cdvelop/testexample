package main

// ContarOcurrencias retorna cuántas veces aparece target en el slice nums.
func ContarOcurrencias(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		if num == target {
			count++
		}
	}
	return count
}
