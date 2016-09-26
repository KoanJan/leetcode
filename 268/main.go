package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	var (
		sum       int = 0
		factorial     = func(n int) int {
			res := 0
			for n > 0 {
				res += n
				n -= 1
			}
			return res
		}
	)
	for _, n := range nums {
		sum += n
	}
	return factorial(len(nums)) - sum
}

func main() {

	// test
	nums := []int{1, 0, 4, 3, 7, 6, 2}
	fmt.Printf("missingNumber(nums) = %d while nums is %v\n", missingNumber(nums), nums)
}
