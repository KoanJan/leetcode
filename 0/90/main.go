package main

import (
	"fmt"
)

func subsetsWithDup(nums []int) [][]int {
	// sort
	simpleSort(nums)
	return swd([]int{}, nums)
}

func swd(pre, nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{cpSlice(pre)}
	}
	res := [][]int{cpSlice(pre)}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := swd(append(pre, nums[i]), nums[i+1:])
		for j := 0; j < len(tmp); j++ {
			res = append(res, tmp[j])
		}
	}
	return res
}

func cpSlice(a []int) []int {
	res := []int{}
	for i := 0; i < len(a); i++ {
		res = append(res, a[i])
	}
	return res
}

func simpleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] > nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
}

func main() {
	// test
	nums := []int{4, 4, 4, 1, 4}
	fmt.Printf("nums: %v\nsubsets: %v\n", nums, subsetsWithDup(nums))
}
