package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if len(obstacleGrid[0]) == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	// mark obstacle with -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}

	// set edges
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == -1 {
			break
		}
		obstacleGrid[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == -1 {
			break
		}
		obstacleGrid[0][j] = 1
	}

	// loop
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			// check if connectable
			if obstacleGrid[i][j] == -1 {
				continue
			}

			a := obstacleGrid[i][j-1]
			b := obstacleGrid[i-1][j]

			if a == -1 && b == -1 {
				obstacleGrid[i][j] = -1
			} else if a == -1 {
				obstacleGrid[i][j] = b
			} else if b == -1 {
				obstacleGrid[i][j] = a
			} else {
				obstacleGrid[i][j] = a + b
			}
		}
	}

	// check if target connectable
	if obstacleGrid[m-1][n-1] == -1 {
		return 0
	}
	return obstacleGrid[m-1][n-1]
}

func main() {
	// test
	obstacleGrid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	fmt.Printf("obstacleGrid: %v\n, unique paths: %d\n", obstacleGrid, uniquePathsWithObstacles(obstacleGrid))
}
