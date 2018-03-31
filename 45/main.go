package main

import "fmt"

// Greed is much better than DP
func jump(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return 0
	}
	count := 0
	rang := 0
	next := 0
	i := 0
	for rang < ln-1 {
		if i <= rang {
			newNext := i + nums[i]
			if next < newNext {
				next = newNext
			}
			i += 1
		} else {
			count += 1
			rang = next
		}
	}
	return count
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
