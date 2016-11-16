package main

import (
	"fmt"
)

func trap(height []int) int {
	var l, r, res int = 0, len(height) - 1, 0
	for l < r {
		if height[l] < height[r] {
			minH := height[l]
			l++
			for l < r && height[l] < minH {
				res += minH - height[l]
				l++
			}
		} else {
			minH := height[r]
			r--
			for l < r && height[r] < minH {
				res += minH - height[r]
				r--
			}
		}
	}
	return res
}

func main() {
	// test
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("height: %v\n", height)
	fmt.Printf("result: %d\n", trap(height))
}
