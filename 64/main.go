package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {

	// invalid grid
	if len(grid) == 0 {
		return -1
	}
	if len(grid[0]) == 0 {
		return -1
	}

	// init
	var (
		m   int = len(grid)
		n   int = len(grid[0])
		sum     = make([][]int, m)
	)
	for i := 0; i < m; i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, 0)
		}
		sum[i] = tmp
	}
	sum[0][0] = grid[0][0]

	// store
	for i := 1; i < m; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		sum[0][i] = sum[0][i-1] + grid[0][i]
	}
	less := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sum[i][j] = grid[i][j] + less(sum[i-1][j], sum[i][j-1])
		}
	}

	return sum[m-1][n-1]
}

func main() {

	// test
	grid := [][]int{
		[]int{1, 2},
		[]int{1, 1},
	}
	fmt.Printf("minPathSum of grid %v is %d\n", grid, minPathSum(grid))
}
