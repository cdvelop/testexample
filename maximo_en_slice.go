package main

// MaximoEnSlice retorna el valor máximo dentro de nums.
// Si el slice está vacío, debe retornar 0.
func MaximoEnSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
