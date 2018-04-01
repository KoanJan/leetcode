package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		panic("Invalid input")
	}
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

func main() {
	// test
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Index of one peak in %v is %d\n", nums, findPeakElement(nums))
}
