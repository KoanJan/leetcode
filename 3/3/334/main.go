package main

import (
	"fmt"
)

func increasingTriplet(nums []int) bool {

	var (
		n       int = len(nums)
		i, j, k int = 0, 1, 2
	)

	// nums is too short
	if n < 3 {
		return false
	}

	// find i and tmp j
	for nums[i] >= nums[j] && i < n-2 && j < n-1 {
		i = j
		j = i + 1
	}
	// if no k existed
	if j >= n-1 {
		return false
	}
	k = j + 1
	// find j and k
	for k < n {
		fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)
		if nums[k] > nums[j] {
			return true
		}
		if nums[k] < nums[j] {
			if k < n-1 {
				if nums[k+1] > nums[k] && nums[k+1] < nums[j] {
					// reset i and j
					i = k
					j = k + 1
					k = k + 2
					continue
				}
			}
			if nums[k] > nums[i] {
				j = k
			}
		}
		k += 1
	}
	return false
}

func main() {

	// test
	nums := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	fmt.Printf("increasingTriplet with nums is %v returns %t\n", nums, increasingTriplet(nums))
}
