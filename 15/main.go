package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	simpleSort(nums)
	var (
		res [][]int = [][]int{}
	)
	for i := 0; i < len(nums)-2; i++ {
		var l, r int = i + 1, len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return uniqSliceSlice(res)
}

func simpleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func uniqSliceSlice(ss [][]int) [][]int {
	res := [][]int{}
	for i := 0; i < len(ss); i++ {
		contained := false
		for j := 0; j < len(res); j++ {
			if len(ss[i]) == len(res[j]) {
				contained = true
				for k := 0; k < len(ss[i]); k++ {
					if ss[i][k] != res[j][k] {
						contained = false
						break
					}
				}
				if contained {
					break
				}
			}
		}
		if contained {
			continue
		}
		res = append(res, ss[i])
	}
	return res
}

func main() {
	// test
	nums := []int{-2, 0, 1, 1, 2}
	fmt.Printf("given array: %v\nsolution: %v\n", nums, threeSum(nums))
}
