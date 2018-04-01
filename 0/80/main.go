package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 3 {
		return length
	}
	var (
		i, j, cut int = 0, 2, 0
	)
	for i < length-2-cut && j < length {
		nums[j-cut] = nums[j]
		if nums[i] == nums[j-cut] {
			cut++
		} else {
			i++
		}
		j++
	}
	return length - cut
}

func main() {
	// test
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Printf("removeDuplicates(nums) = %d while nums is %v\n", removeDuplicates(nums), nums)
}
