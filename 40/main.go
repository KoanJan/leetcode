package main

import (
	"fmt"
)

func combinationSum2(candidates []int, target int) [][]int {
	return uniqSlice(cs2(candidates, target, []int{}))
}

func cs2(candidates []int, target int, pre []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		if candidates[i] == target {
			res = append(res, append(cpSlice(pre), target))
		} else {
			subRes := cs2(candidates[i+1:], target-candidates[i], append(cpSlice(pre), candidates[i]))
			for j := 0; j < len(subRes); j++ {
				res = append(res, subRes[j])
			}
		}
	}
	return res
}

func cpSlice(a []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i]
	}
	return res
}

func uniqSlice(a [][]int) [][]int {
	res := [][]int{}
	for i := 0; i < len(a); i++ {
		hasContained := false
		for j := 0; j < len(res); j++ {
			if len(a[i]) == len(res[j]) {

				equal := true
				ma := map[int]int{}
				for _, e := range a[i] {
					ma[e] += 1
				}
				mr := map[int]int{}
				for _, e := range res[j] {
					mr[e] += 1
				}
				for k, v := range ma {
					if mr[k] != v {
						equal = false
						break
					}
				}
				if equal {
					hasContained = true
					break
				}
			}
			if hasContained {
				break
			}
		}
		if !hasContained {
			res = append(res, a[i])
		}
	}
	return res
}

func main() {
	// test
	var (
		candidates []int = []int{10, 1, 2, 7, 6, 1, 5}
		target     int   = 8
	)
	fmt.Printf("candidates: %v\ntarget: %d\n\ncombinations: %v\n", candidates, target, combinationSum2(candidates, target))
}
