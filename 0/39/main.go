package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	return cs(candidates, target, []int{})
}

func cs(candidates []int, target int, preRes []int) [][]int {
	if target == 0 {
		return [][]int{cpSlice(preRes)}
	}
	res := [][]int{}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] <= target {
			tmp := cs(candidates[i:], target-candidates[i], append(preRes, candidates[i]))
			for j := 0; j < len(tmp); j++ {
				res = append(res, tmp[j])
			}
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

func main() {
	// test
	var (
		candidates []int = []int{2, 3, 6, 7, 8}
		target     int   = 15
	)
	fmt.Printf("condidates: %v\ntarget: %d\n\ncombinationSum: %v\n", candidates, target, combinationSum(candidates, target))
}
