package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	l, r := bs(nums, 0, len(nums)-1, target)
	if nums[l] > target {
		return 0
	} else if nums[l] == target {
		return l
	} else if nums[r] >= target {
		return r
	} else {
		return r + 1
	}
}

func bs(nums []int, l, r, target int) (int, int) {
	if r-l <= 1 {
		return l, r
	}
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid, mid
	} else if nums[mid] < target {
		return bs(nums, mid, r, target)
	} else {
		return bs(nums, l, mid, target)
	}
}

func main() {
	// test
	var (
		nums   = []int{1, 3}
		target = 0
	)
	fmt.Printf("searchInsert with nums is %v and target is %d makes result %d\n", nums, target, searchInsert(nums, target))
}
