package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	var (
		data map[int]bool = make(map[int]bool)
		max  int          = 0
	)
	for i := 0; i < len(nums); i++ {
		data[nums[i]] = true
	}
	for i := 0; i < len(nums); i++ {
		if data[nums[i]] {
			delete(data, nums[i])
			count := 1
			for x := nums[i] + 1; data[x]; x++ {
				count++
				delete(data, x)
			}
			for x := nums[i] - 1; data[x]; x-- {
				count++
				delete(data, x)
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}

func main() {
	// test
	nums := []int{100, 4, 200, 1, 3, 2, 2}
	fmt.Printf("The length of longest consecutive subsequence of %v is %d\n", nums, longestConsecutive(nums))
}
