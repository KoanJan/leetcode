package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		newNums := []int{}
		for j := 0; j < len(nums); j++ {
			if j != i {
				newNums = append(newNums, nums[j])
			}
		}
		subRes := permute(newNums)
		for _, sr := range subRes {

			// it can be like this, however, leetcode requires f**king order of results
			/*res = append(res, append(sr, nums[i]))*/

			r := []int{nums[i]}
			for _, e := range sr {
				r = append(r, e)
			}
			res = append(res, r)
		}
	}
	return res
}

func main() {
	// test
	nums := []int{1, 2, 3}
	fmt.Printf("nums is %v, and permutations are \n%v\n", nums, permute(nums))
}
