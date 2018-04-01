package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	// init
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	var (
		i, j    int = 0, -1
		steps   int = 1
		forward int = 4 // up: 1, down: 2, left: 3, right: 4
	)
	// loop
	for steps <= n*n {
		switch forward {
		// up
		case 1:
			if i == 0 || res[i-1][j] != 0 {
				// change forward right
				forward = 4
				j++
			} else {
				i--
			}
		// down
		case 2:
			if i == n-1 || res[i+1][j] != 0 {
				// change forward left
				forward = 3
				j--
			} else {
				i++
			}
		// left
		case 3:
			if j == 0 || res[i][j-1] != 0 {
				// change forward up
				forward = 1
				i--
			} else {
				j--
			}
		// right
		case 4:
			if j == n-1 || res[i][j+1] != 0 {
				// change forward down
				forward = 2
				i++
			} else {
				j++
			}
		}
		// go
		res[i][j] = steps
		steps++
	}
	return res
}

func main() {
	// test
	n := 0
	fmt.Printf("generateMatrix(%d) gets %v\n", n, generateMatrix(n))
}
