package main

import (
	"fmt"
)

func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	res := 0
	// it should be like it but time out
	/*for i := 0; i < n; i++ {
		res += numTrees(i) * numTrees(n-1-i)
	}*/
	if n <= 2 {
		for i := 0; i < n; i++ {
			res += numTrees(i) * numTrees(n-1-i)
		}
	} else {
		for i := 0; i < n/2; i++ {
			res += 2 * numTrees(i) * numTrees(n-1-i)
		}
		if n%2 == 1 {
			res += numTrees((n-1)/2) * numTrees((n-1)/2)
		}
	}
	return res
}

func main() {
	// test
	n := 5
	fmt.Printf("numTrees(%d) = %d\n", n, numTrees(n))
}
