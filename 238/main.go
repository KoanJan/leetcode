package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	var (
		n   int   = len(nums)
		res []int = make([]int, n)
		p   int   = 1
	)
	// init
	res[0] = 1
	res[n-1] = 1
	// left loop
	for i := 1; i < n; i++ {
		p *= nums[i-1]
		res[i] = p
	}
	// right loop
	p = 1 // reset p
	for i := n - 2; i >= 0; i-- {
		p *= nums[i+1]
		res[i] *= p
	}
	// return
	return res
}

func main() {
	// test
	nums := []int{9, 0, -2}
	fmt.Printf("productExceptSelf(nums) = %v while nums is %v", productExceptSelf(nums), nums)
}
