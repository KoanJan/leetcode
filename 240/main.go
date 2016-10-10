package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	if m == 0 {
		return false
	}

	// find x,y
	var (
		x, y int
	)
	mA := []int{}
	for _, mE := range matrix {
		mA = append(mA, mE[0])
	}
	x1, x2 := scan(mA, target)
	if matrix[x2][0] == target {
		return true
	}
	if matrix[x2][0] < target {
		x = x2
	} else {
		x = x1
	}
	y1, y2 := scan(matrix[0], target)
	if matrix[0][y2] == target {
		return true
	}
	if matrix[0][y2] < target {
		y = y2
	} else {
		y = y1
	}

	// scan
	for i := 0; i <= x; i++ {
		a, b := scan(matrix[i][0:y+1], target)
		if matrix[i][a] == target || matrix[i][b] == target {
			return true
		}
	}
	return false
}

func scan(nums []int, target int) (int, int) {
	if len(nums) == 0 {
		panic("Invalid input")
	}
	var (
		l   int = 0
		r   int = len(nums) - 1
		mid int = (l + r) / 2
	)
	for r-l > 1 {
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			return mid, mid
		}
		mid = (l + r) / 2
	}
	return l, r
}

func main() {

	// test
	var (
		matrix = [][]int{
			[]int{1, 4, 7, 11, 15},
			[]int{2, 5, 8, 12, 19},
			[]int{3, 6, 9, 16, 22},
			[]int{10, 13, 14, 17, 24},
			[]int{18, 21, 23, 26, 30},
		}
		target = 5
	)
	fmt.Printf("search %d in matrix %v, got %t\n", target, matrix, searchMatrix(matrix, target))
}
