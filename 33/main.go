package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	var (
		l, r, minIdx int = 0, len(nums) - 1, -1
		mid          int = (l + r) / 2
	)

	// find minIdx
	for r-l > 1 {
		if nums[mid] > nums[r] {
			l = mid
		} else {
			r = mid
		}
		mid = (l + r) / 2
	}
	if nums[r] > nums[l] {
		minIdx = l
	} else {
		minIdx = r
	}

	// find target
	if minIdx > 0 {
		if target > nums[len(nums)-1] {
			l, r = 0, minIdx-1
			mid = (l + r) / 2
		} else {
			l, r = minIdx, len(nums)-1
			mid = (l + r) / 2
		}
	} else {
		l, r = 0, len(nums)-1
	}
	for r-l > 1 {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
		mid = (l + r) / 2
	}
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}

func main() {
	// test
	var (
		nums   []int = []int{1}
		target int   = 2
	)
	fmt.Printf("search %d in %v, got %d\n", target, nums, search(nums, target))
}
