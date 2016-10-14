package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		panic("Invalid input")
	}
	// init
	var (
		extraRow int = -1
		extraCol int = -1
	)
	// first column
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			extraCol = 0
			break
		}
	}
	// first row
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			extraRow = 0
			break
		}
	}
	// store
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// set
	// 1.set row
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	// 2.set column
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	// 3.set first column and row
	if extraRow == 0 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if extraCol == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

func matrixString(matrix [][]int) string {
	res := ""
	for _, row := range matrix {
		for _, e := range row {
			res += fmt.Sprintf(" %2d", e)
		}
		res += "\n"
	}
	return res
}

func main() {
	// test
	matrix := [][]int{
		[]int{0, 0, 0, 5},
		[]int{4, 3, 1, 4},
		[]int{0, 1, 1, 4},
		[]int{1, 2, 1, 3},
		[]int{0, 0, 1, 1},
	}
	fmt.Printf("before: \n%s\n", matrixString(matrix))
	setZeroes(matrix)
	fmt.Printf("after: \n%s\n", matrixString(matrix))
}
