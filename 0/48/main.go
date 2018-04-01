package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	var (
		n = len(matrix)
	)
	// rotate
	if n <= 1 {
		return
	}
	for t := 0; t <= n/2; t++ {
		l := n - 2*t
		fmt.Printf("l: %d\n", l)
		for i := 0; i < l-1; i++ {
			fmt.Printf("swapping...\n")

			tmp := matrix[i+(n-l)/2][(n-l)/2]

			fmt.Printf("swapping (%d, %d) and (%d, %d)\n", i+(n-l)/2, (n-l)/2, l-1+(n-l)/2, i+(n-l)/2)
			matrix[i+(n-l)/2][(n-l)/2] = matrix[l-1+(n-l)/2][i+(n-l)/2]

			fmt.Printf("swapping (%d, %d) and (%d, %d)\n", l-1+(n-l)/2, i+(n-l)/2, l-1-i+(n-l)/2, l-1+(n-l)/2)
			matrix[l-1+(n-l)/2][i+(n-l)/2] = matrix[l-1-i+(n-l)/2][l-1+(n-l)/2]

			fmt.Printf("swapping (%d, %d) and (%d, %d)\n", l-1-i+(n-l)/2, l-1+(n-l)/2, (n-l)/2, l-1-i+(n-l)/2)
			matrix[l-1-i+(n-l)/2][l-1+(n-l)/2] = matrix[(n-l)/2][l-1-i+(n-l)/2]

			fmt.Printf("swapping (%d, %d) and (%d, %d)\n", (n-l)/2, l-1-i+(n-l)/2, i+(n-l)/2, (n-l)/2)
			matrix[(n-l)/2][l-1-i+(n-l)/2] = tmp

			fmt.Printf("swapping matrix: \n%s\n", MatrixString(matrix))
		}
	}
}

func MatrixString(matrix [][]int) string {
	res := ""
	for _, m := range matrix {
		for _, e := range m {
			res += fmt.Sprintf(" %2d", e)
		}
		res += "\n"
	}
	res += "\n"
	return res
}

func main() {

	// test
	matrix := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15},
		[]int{16, 17, 18, 19, 20},
		[]int{21, 22, 23, 24, 25},
	}
	// matrix = [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// }
	// matrix = [][]int{
	// 	[]int{1, 2},
	// 	[]int{3, 4},
	// }
	fmt.Printf("matrix: \n%s\n", MatrixString(matrix))
	rotate(matrix)
	fmt.Printf("matrix: \n%s\n", MatrixString(matrix))
}
