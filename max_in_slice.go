package wordcount

func MaxInSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
