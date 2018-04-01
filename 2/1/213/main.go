package main

import (
	"fmt"
)

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	var (
		max = make([]int, n+1)
		m1  int
		m2  int
	)

	// except first house
	max[2] = nums[1]
	for i := 2; i < n+1; i++ {
		max[i] = max[i-2] + nums[i-1]
		if max[i] < max[i-1] {
			max[i] = max[i-1]
		}
	}
	m1 = max[n]

	// except last house
	max = make([]int, n+1)
	max[2] = nums[0]
	for i := 2; i < n+1; i++ {
		max[i] = max[i-2] + nums[i-2]
		if max[i] < max[i-1] {
			max[i] = max[i-1]
		}
	}
	m2 = max[n]

	// compare
	if m1 > m2 {
		return m1
	}
	return m2
}

func main() {
	// test
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 5, 6, 7, 8, 9}
	fmt.Printf("Max amount by rob of houses %v : %d\n", nums, rob(nums))
}
