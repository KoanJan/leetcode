package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {

	// special
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}
	if m == 1 {
		return matrix[0]
	}
	if n == 1 {
		tmp := make([]int, m)
		for i := 0; i < m; i++ {
			tmp[i] = matrix[i][0]
		}
		return tmp
	}

	// common
	var (
		res        []int    = make([]int, m*n)
		track      [][]bool = make([][]bool, m)
		i, j, c, d int      = 0, 0, 0, 1
	)
	for i := 0; i < m; i++ {
		track[i] = make([]bool, n)
	}
	for c < m*n {

		res[c] = matrix[i][j]
		track[i][j] = true
		c++
		switch d {
		// Right
		case 1:
			if j == n-1 || track[i][j+1] {
				i++
				d = 2
			} else {
				j++
			}
		// Down
		case 2:
			if i == m-1 || track[i+1][j] {
				j--
				d = 3
			} else {
				i++
			}
		// Left
		case 3:
			if j == 0 || track[i][j-1] {
				i--
				d = 4
			} else {
				j--
			}
		// Up
		case 4:
			if i == 0 || track[i-1][j] {
				j++
				d = 1
			} else {
				i--
			}
		}
	}

	return res
}

func main() {
	// test
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Printf("matrix: %v\nspiralOrder: %v\n", matrix, spiralOrder(matrix))
}
