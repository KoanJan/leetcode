package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var (
		n   int = len(height)
		max int = 0
	)
	for i := n - 1; i > 0; i-- {
		for j := 0; j+i < n; j++ {
			var (
				h, area int
			)
			if height[j] < height[j+i] {
				h = height[j]
			} else {
				h = height[j+i]
			}
			area = h * i
			if area > max {
				max = area
			}
		}
	}
	return max
}

func main() {
	// test
	height := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("height is %v, max area is %d\n", height, maxArea(height))
}
