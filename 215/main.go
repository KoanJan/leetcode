package main

import (
	"fmt"
)

// func findKthLargest(nums []int, k int) int {
// 	var (
// 		pivot      int = nums[0]
// 		pivotCount int = 1
// 		larger         = []int{}
// 		smaller        = []int{}
// 	)
// 	// divive
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] < pivot {
// 			smaller = append(smaller, nums[i])
// 		} else if nums[i] > pivot {
// 			larger = append(larger, nums[i])
// 		} else {
// 			pivotCount++
// 		}
// 	}
// 	fmt.Printf("pivot is %d and k is %d\n", pivot, k)
// 	fmt.Printf("larger: %v\nsmaller: %v\n", larger, smaller)
// 	if len(larger) >= k {

// 		fmt.Println("k in larger\n")
// 		return findKthLargest(larger, k)

// 	} else if len(nums)-k < len(smaller) {

// 		fmt.Println("k in smaller\n")
// 		return findKthLargest(smaller, k-len(larger)-pivotCount)

// 	} else {
// 		fmt.Println("found k\n")
// 		return pivot
// 	}
// }

func findKthLargest(nums []int, k int) int {
	fmt.Printf("\nnums: %v\n", nums)
	if len(nums) == 1 && k == 1 {
		return nums[0]
	}
	var (
		pivot      int = nums[0]
		pivotCount int = 1
		l          int = 1
		r          int = len(nums) - 1
	)
	// divive
	for l < r {
		for l < len(nums) {
			if nums[l] > pivot {
				l++
			} else if nums[l] == pivot {
				tmp := nums[l]
				nums[l] = nums[pivotCount]
				nums[pivotCount] = tmp
				pivotCount++
			} else {
				break
			}
		}
		for r > 0 {
			if nums[r] < pivot {
				r--
			} else if nums[r] == pivot {
				tmp := nums[r]
				nums[r] = nums[pivotCount]
				nums[pivotCount] = tmp
				pivotCount++
			} else {
				break
			}
		}
		// swap
		if l < len(nums) && r > 0 {
			tmp := nums[l]
			nums[l] = nums[r]
			nums[r] = tmp
		}
	}
	fmt.Printf("pivot: %d\n", pivot)
	fmt.Printf("pivotCount: %d\n", pivotCount)
	fmt.Printf("k: %d\n", k)

	// last time
	l--
	if r != len(nums)-1 {
		r++
	}
	fmt.Printf("l: %d, r: %d\n", l, r)
	if nums[l] < nums[r] {

		tmp := nums[l]
		nums[l] = nums[r]
		nums[r] = tmp
	}
	fmt.Printf("l-r=%d\n", l-r)

	fmt.Printf("nums: %v\n", nums)
	if l-pivotCount+1 >= k {

		fmt.Printf("k in larger\n")
		fmt.Printf("larger: %v\n", nums[pivotCount:l+1])
		return findKthLargest(nums[pivotCount:l+1], k)

	} else if k > r {

		fmt.Printf("k in smaller\n")
		fmt.Printf("smaller: %v\n", nums[r:])
		return findKthLargest(nums[r:], k-r)

	} else {
		fmt.Println("found k\n")
		return pivot
	}
}

func main() {
	// test
	var (
		nums     = []int{1, 4, 3, 2, 9, 8, 5, 7, 6}
		k    int = 7
	)
	fmt.Printf("find %dth Largest in %v got %d\n", k, nums, findKthLargest(nums, k))
}
