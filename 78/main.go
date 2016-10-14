package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	res := genSets([]int{}, nums)
	for i := 0; i < len(res); i++ {
		fuckingSort(res[i])
	}
	return res
}

func genSets(pre []int, nums []int) [][]int {
	res := [][]int{pre}
	for i := 0; i < len(nums); i++ {
		sets := genSets(append(cpSlice(pre), nums[i]), nums[0:i])
		for _, r := range sets {
			res = append(res, r)
		}
	}
	return res
}

func cpSlice(src []int) []int {
	res := []int{}
	for i := 0; i < len(src); i++ {
		res = append(res, src[i])
	}
	return res
}

func fuckingSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[i] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
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
	nums := []int{9, 0, 3, 5, 7}
	fmt.Printf("subsets of %v:\n%s", nums, subSetString(subsets(nums)))
}
