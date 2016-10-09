package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	var (
		max int = 0
		sum int = 0
	)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < 0 {
			sum = 0
		}
		if sum > max {
			max = sum
		}
	}
	// all numbers are negetive
	if max == 0 {
		max = nums[0]
		for i := 0; i < len(nums); i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
	}
	return max
}

func main() {
	// test
	nums := []int{-2, -1, -3, -4, -1, -2, -1, -5, -4}
	fmt.Printf("While nums is %v, maxSubArray returns %d\n", nums, maxSubArray(nums))
}
