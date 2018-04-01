package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	var (
		l1, r1 = 0, len(nums) - 1
		mid1   = (r1 + l1) / 2

		l2, r2 = 0, len(nums) - 1
		mid2   = (r2 + l2) / 2

		l, r int = -1, -1
	)
	// left
	for r1-l1 > 1 {
		if nums[mid1] < target {
			l1 = mid1
		} else {
			r1 = mid1
		}
		mid1 = (l1 + r1) / 2
	}
	if nums[l1] == target {
		l = l1
	} else if nums[r1] == target {
		l = r1
	}
	// right
	for r2-l2 > 1 {
		if nums[mid2] > target {
			r2 = mid2
		} else {
			l2 = mid2
		}
		mid2 = (l2 + r2) / 2
	}
	if nums[r2] == target {
		r = r2
	} else if nums[l2] == target {
		r = l2
	}

	return []int{l, r}
}

func main() {
	// test
	var (
		nums   []int = []int{}
		target int   = 8
	)
	fmt.Printf("Search %d in range %v got %v\n", target, nums, searchRange(nums, target))
}
