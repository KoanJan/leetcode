package main

import (
	"fmt"
)

func numSquares(n int) int {

	// init
	data := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		data[i] = n + 1
	}
	for i := 1; i*i <= n; i++ {
		data[i*i] = 1
	}

	// find
	for a := 0; a <= n; a++ {
		for b := 0; a+b*b <= n; b++ {
			if data[a]+1 < data[a+b*b] {
				data[a+b*b] = data[a] + 1
			}
		}
	}

	return data[n]
}

func main() {
	// test
	n := 999999
	fmt.Printf("numSquares(%d) = %d\n", n, numSquares(n))
}
