package testexample

import (
	"fmt"
)

func SumPositives(nums []int) int {
	sum := 0
	for _, n := range nums {
		if n > 0 {
			sum += n
		}
	}
	return sum
}

func slice() {

	result := SumPositives([]int{1, 2, 3})
	fmt.Println(result)
}
