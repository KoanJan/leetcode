package main

import (
	"fmt"
)

func findMin(nums []int) int {
	if len(nums) == 0 {
		panic("Invalid input")
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var (
		l   int = 0
		r   int = len(nums) - 1
		mid int = (l + r) / 2
	)
	for r-l > 1 {
		if nums[mid] > nums[r] {
			l = mid
		} else if nums[mid] < nums[l] {
			r = mid
		} else {
			return nums[0]
		}
		mid = (l + r) / 2
	}
	if nums[l] < nums[r] {
		return nums[l]
	}
	return nums[r]
}

func main() {
	// test
	nums := []int{1, 2, 3, -3, -2}
	fmt.Printf("findMin with %v got %d\n", nums, findMin(nums))
}
