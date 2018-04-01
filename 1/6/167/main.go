package main

import (
	"fmt"
)

// length of numbers must be 2 at least
func twoSum(numbers []int, target int) []int {

	// find target
	var (
		n   int = len(numbers)
		l   int = 0
		r   int = n
		mid int = (l + r) / 2
		k   int = mid
	)
	for l < r-1 {
		if numbers[mid] < target {
			l = mid
		} else if numbers[mid] > target {
			r = mid
		} else {
			break
		}
		k = mid
		mid = (l + r) / 2
	}
	if numbers[k] == target {
		k -= 1
	}

	// now the searching field is 0 to k
	var (
		a int = 0
		b int = k
	)
	for a < b {
		if numbers[a]+numbers[b] > target {
			b--
		} else if numbers[a]+numbers[b] < target {
			a++
		} else {
			// equal, find out
			return []int{a + 1, b + 1}
		}
	}

	// find no results
	return []int{-1, -1}
}

func main() {

	// test
	var (
		nums   = []int{1, 2, 5, 7, 9}
		target = 9
	)
	fmt.Printf("twoSum(nums, target) = %v while sums is %v and target is %d\n", twoSum(nums, target), nums, target)
}
