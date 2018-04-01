package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	var l, r, mid, row int

	// row
	l = 0
	r = m - 1
	mid = (l + r) / 2
	for l < r-1 {
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target {
			r = mid
		} else {
			l = mid
		}
		mid = (l + r) / 2
	}
	if matrix[r][0] == target {
		return true
	} else if matrix[r][0] < target {
		row = r
	} else {
		row = l
	}

	// col
	l = 0
	r = n - 1
	mid = (l + r) / 2
	for l < r-1 {
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] > target {
			r = mid
		} else {
			l = mid
		}
		mid = (l + r) / 2
	}
	return matrix[row][r] == target || matrix[row][l] == target
}

func main() {
	// test
	var (
		matrix = [][]int{
			[]int{1, 3, 5, 7},
			[]int{10, 11, 16, 20},
			[]int{23, 30, 34, 50},
		}
		target int = 50
	)
	fmt.Printf("%d in matrix %v: %t\n", target, matrix, searchMatrix(matrix, target))
}
