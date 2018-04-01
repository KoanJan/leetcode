package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var (
		reachable []bool = make([]bool, len(nums))
		n         int    = len(nums)
	)
	reachable[0] = true
	for i := 0; i < n; i++ {
		if !reachable[i] {
			return false
		}
		for j := i; j <= i+nums[i] && j < n; j++ {
			reachable[j] = true
		}
	}
	return true
}

func main() {
	// test
	nums := []int{1, 2}
	fmt.Printf("nums: %v\ncanJump? %t\n", nums, canJump(nums))
}
