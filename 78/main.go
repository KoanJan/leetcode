package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	return genSets([]int{}, nums)
}

func genSets(pre []int, nums []int) [][]int {
	res := [][]int{pre}
	for i := 0; i < len(nums); i++ {
		newPre := append(pre, nums[i])
		sets := genSets(newPre, nums[i+1:])
		for _, r := range sets {
			res = append(res, r)
		}
	}
	return res
}

func subSetString(set [][]int) string {
	res := ""
	for _, s := range set {
		for _, e := range s {
			res += fmt.Sprintf(" %2d", e)
		}
		res += "\n"
	}
	return res
}

func main() {
	// test
	nums := []int{1, 2, 3}
	fmt.Printf("subsets of %v:\n%s", nums, subSetString(subsets(nums)))
}
