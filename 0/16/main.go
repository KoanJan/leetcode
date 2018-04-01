package main

import (
	"fmt"
)

func threeSumClosest(nums []int, target int) int {

	nums = mergeSort(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		var l, r int = i + 1, len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			} else if sum > target {
				r--
			} else {
				l++
			}
			if absDiff(closest, target) > absDiff(sum, target) {
				closest = sum
			}
		}
	}

	return closest
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// mergeSort
func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	return merge(mergeSort(a[0:len(a)/2]), mergeSort(a[len(a)/2:]))
}

func merge(a []int, b []int) []int {
	var (
		res []int = []int{}
		ca  int   = 0
		cb  int   = 0
	)
	for {
		// index out
		if ca == len(a) || cb == len(b) {
			break
		}
		if a[ca] < b[cb] {
			res = append(res, a[ca])
			ca++
		} else {
			res = append(res, b[cb])
			cb++
		}
	}
	for ca < len(a) {
		res = append(res, a[ca])
		ca++
	}
	for cb < len(b) {
		res = append(res, b[cb])
		cb++
	}
	return res
}

func main() {
	// test
	var (
		nums   []int = []int{0, 1, 2}
		target int   = 0
	)
	fmt.Printf("nums: %v\ntarget: %d\n\nthreeSumClosest: %d\n", nums, target, threeSumClosest(nums, target))
}
