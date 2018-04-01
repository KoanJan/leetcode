package main

import (
	"fmt"
)

func hIndex(citations []int) int {
	n := len(citations)
	for i := n - 1; i >= 0; i-- {
		if citations[i] < n-i {
			return n - i - 1
		}
	}
	return n
}

func main() {
	// test
	citations := []int{2, 4, 4, 5, 7, 7, 7, 7, 8}
	fmt.Printf("H-Index of ascending %v is %d\n", citations, hIndex(citations))
}
