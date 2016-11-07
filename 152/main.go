package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		products []int = make([]int, len(nums))
		max      int   = nums[0]
	)
	products[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		var tmp, m int = nums[i], products[i-1]
		if tmp > m {
			m = tmp
		}
		for j := i - 1; j >= 0; j-- {
			tmp *= nums[j]
			if tmp > m {
				m = tmp
			}
		}
		if m > max {
			max = m
		}
	}

	return max
}

func main() {
	// test
	nums := []int{2, 3, -2, 4}
	fmt.Printf("Maximum product of %v is %d\n", nums, maxProduct(nums))
}
