package main

import (
	"fmt"
)

func majorityElement(nums []int) []int {
	res := []int{}
	if len(nums) == 0 {
		return res
	}
	var (
		n1, n2         int = nums[0], 0
		count1, count2 int = 1, 0
	)
	// find out most frequent number
	for i := 1; i < len(nums); i++ {
		if nums[i] == n1 {
			count1++
		} else if nums[i] == n2 {
			count2++
		} else {
			if count1 == 0 {
				n1 = nums[i]
				count1++
			} else if count2 == 0 {
				n2 = nums[i]
				count2++
			} else {
				count1--
				count2--
			}
		}
	}
	// reset count
	count1 = 0
	count2 = 0
	// count
	for i := 0; i < len(nums); i++ {
		if nums[i] == n1 {
			count1++
		}
		if nums[i] == n2 {
			count2++
		}
	}
	// check count if larger than length/3
	if count1 > len(nums)/3 {
		res = append(res, n1)
	}
	if n1 == n2 {
		return res
	}
	if count2 > len(nums)/3 {
		res = append(res, n2)
	}
	return res
}

func main() {
	// test
	nums := []int{}
	fmt.Printf("nums: %v\nmajorityElements: %v\n", nums, majorityElement(nums))
}
