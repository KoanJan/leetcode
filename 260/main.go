package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {
	// calc xor
	var (
		xor int = 0
		a   int = 0
		b   int = 0
	)
	for _, num := range nums {
		xor ^= num
	}
	lb := xor & -xor
	for _, num := range nums {
		if num&lb != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

func main() {

	// test
	nums := []int{1, 5, 3, 5, 1, 6, 7, 8, 8, 9, 2, 4, 2, 9, 3, 4}
	fmt.Printf("singleNumber(nums) = %v while nums is %v", singleNumber(nums), nums)
}
