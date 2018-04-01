package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	return cb(1, n, k)
}

func cb(s, n int, k int) [][]int {
	if k == 0 {
		return make([][]int, 1)
	}
	// fmt.Printf("s: %d, n: %d, k: %d\n", s, n, k)
	res := [][]int{}
	for i := s; i <= n+1-k; i++ {
		subRes := cb(i+1, n, k-1)
		for _, subSeq := range subRes {
			t := []int{i}
			for _, e := range subSeq {
				t = append(t, e)
			}
			res = append(res, t)
		}
	}
	return res
}

func main() {
	// test
	var (
		n = 10
		k = 5
	)
	fmt.Printf("combine(%d, %d) got %v\n", n, k, combine(n, k))
}
