package main

import (
	"fmt"
)

func combinationSum3(k int, n int) [][]int {
	return cs3(k, n, 1, []int{})
}

func cs3(k, n, s int, base []int) [][]int {

	if k == 0 && n == 0 {
		return [][]int{base}
	}

	res := [][]int{}
	if k > 0 && n > 0 {
		for i := s; i <= 9 && i <= n; i++ {

			// newBase := append(base, i)
			newBase := []int{}
			for _, b := range base {
				newBase = append(newBase, b)
			}
			newBase = append(newBase, i)

			tmp := cs3(k-1, n-i, i+1, newBase)
			if len(tmp) > 0 {
				for _, c := range tmp {
					res = append(res, c)
				}
			}
		}
	}
	return res
}

func main() {
	// test
	var (
		k int = 8
		n int = 36
	)
	fmt.Printf("combinationSum3(%d, %d):\nresult: %v\n", k, n, combinationSum3(k, n))
}
