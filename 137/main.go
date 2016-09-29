package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var (
		res int = 0
		cur int = 1
	)
	// 32-bit
	for i := 0; i < 32; i++ {
		times := 0
		for _, n := range nums {
			if n&cur > 0 {
				times++
			}
		}
		// it won't be 2
		if times%3 == 1 {
			res |= cur
		}
		cur = 1 << uint(i+1)
	}
	return res
}

func main() {

	// test
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	fmt.Printf("In %v, single num is %d\n", nums, singleNumber(nums))
}
