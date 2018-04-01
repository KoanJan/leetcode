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
	)
	// 32-bit
	for i := 0; i < 32; i++ {
		var (
			times int = 0
			cur       = 1 << uint(i)
		)
		for _, n := range nums {
			if n&cur > 0 {
				times++
			}
		}
		// it won't be 2
		if times%3 == 1 {
			res |= cur
		}
	}
	// int is 32 or 64 bit, which depends on OS
	return int(int32(res))
}

func main() {

	// test
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	// nums := []int{1, 1, 1, 4}
	fmt.Printf("In %v, single num is %d\n", nums, singleNumber(nums))
}
