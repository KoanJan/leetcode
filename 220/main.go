package main

import (
	"fmt"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if abs(i-j) <= k && abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// test
	var (
		nums []int = []int{1, 2, 1}
		k, t int   = 1, 1
	)
	fmt.Printf("nums: %v\nk: %d\nt: %d\ncontainsNearbyAlmostDuplicate: %t\n", nums, k, t, containsNearbyAlmostDuplicate(nums, k, t))
}
