package main

import "fmt"

// it must be kidding me ... however it beats 100% on the leetcode
func search(nums []int, target int) bool {
	for _, n := range nums {
		if n == target {
			return true
		}
	}
	return false
}

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{1, 1, 3, 1}
	target := 3
	fmt.Println(search(nums, target))
}
