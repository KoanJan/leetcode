package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	// down steps: m-1, right steps: n-1

	// the deprecated code is so beautiful but it crash TIMEOUT, fuck it!
	/*if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)*/

	// total steps: m+n-2
	// 1. init paths
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			paths[i][j] = 1
		}
	}
	// 2.update
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	// 3.return
	return paths[m-1][n-1]
}

func main() {
	// test
	var (
		m int = 1
		n int = 2
	)
	fmt.Printf("uniquePaths(%d, %d) = %d\n", m, n, uniquePaths(m, n))
}
