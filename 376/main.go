package main

import (
	"fmt"
)

// func wiggleMaxLength(nums []int) int {
// 	a := wml(nums, true)
// 	b := wml(nums, false)
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func wml(nums []int, side bool) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	if len(nums) == 1 {
// 		return 1
// 	}
// 	// init
// 	var max int = 1

// 	for i := 1; i < len(nums); i++ {
// 		if (side && nums[i] > nums[0]) || (!side && nums[i] < nums[0]) {
// 			fmt.Printf("nums: %v, side: %t\n", nums, side)
// 			tmp := 1 + wml(nums[i:], !side)
// 			if tmp > max {
// 				max = tmp
// 			}
// 			if max >= len(nums)-i {
// 				break
// 			}
// 		}
// 	}
// 	fmt.Printf("max: %d\n", max)
// 	return max
// }

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	// init
	var (
		res  int = 1
		side bool
		k1   int = 0
		k2   int = 1
	)
	// side init
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[0] {
			continue
		}
		side = nums[i] > nums[0]
		k1 = i
		k2 = i + 1
		res++
		break
	}
	for k2 < len(nums) {
		if (side && nums[k1] > nums[k2]) || ((!side) && nums[k1] < nums[k2]) {
			res++
			side = !side
		}
		k1++
		k2++
	}
	return res
}

func main() {
	// test
	nums := []int{33, 53, 12, 64, 50, 41, 45, 21, 97, 35, 47, 92, 39, 0, 93, 55, 40, 46, 69, 42, 6, 95, 51, 68, 72, 9, 32, 84, 34, 64, 6, 2, 26, 98, 3, 43, 30, 60, 3, 68, 82, 9, 97, 19, 27, 98, 99, 4, 30, 96, 37, 9, 78, 43, 64, 4, 65, 30, 84, 90, 87, 64, 18, 50, 60, 1, 40, 32, 48, 50, 76, 100, 57, 29, 63, 53, 46, 57, 93, 98, 42, 80, 82, 9, 41, 55, 69, 84, 82, 79, 30, 79, 18, 97, 67, 23, 52, 38, 74, 15}
	// nums := []int{33, 53, 12, 64, 50, 41, 45, 21, 97, 35, 47, 92, 39, 0, 93, 55, 40, 46, 69, 42, 6, 95, 51, 68, 72, 9, 32, 84, 34, 64, 6, 2, 26, 98, 3}
	fmt.Printf("while nums is %v, wiggleMaxLength(nums) got %d\n", nums, wiggleMaxLength(nums))
}
