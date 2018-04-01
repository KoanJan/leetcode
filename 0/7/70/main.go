package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	tmp, a, b := 1, 1, 2
	for ; n > 2; n-- {
		tmp = a
		a = b
		b = tmp + a
	}
	return b
}

func main() {
	n := 44
	fmt.Println(climbStairs(n))
}
