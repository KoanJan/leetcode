package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	var row = len(triangle)

	for i := 1; i < row; i++ {

		triangle[i][0] = triangle[i-1][0] + triangle[i][0]

		var j int
		for j = 1; j < len(triangle[i])-1; j++ {
			min1 := triangle[i][j] + triangle[i-1][j-1]
			min2 := triangle[i][j] + triangle[i-1][j]
			if min1 < min2 {
				triangle[i][j] = min1
			} else {
				triangle[i][j] = min2
			}
		}

		triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
	}

	min := triangle[row-1][0]
	for i := 1; i < len(triangle[row-1]); i++ {
		if triangle[row-1][i] < min {
			min = triangle[row-1][i]
		}
	}
	return min
}

func main() {
	// test
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	fmt.Printf("triangle: %v\n", triangle)
	fmt.Printf("minimumTotal = %d\n", minimumTotal(triangle))
}
