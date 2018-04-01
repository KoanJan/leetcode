package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var res *TreeNode
	switch len(nums) {
	case 0:
	case 1:
		res = &TreeNode{Val: nums[0]}
	case 2:
		res = &TreeNode{Val: nums[1]}
		res.Left = &TreeNode{Val: nums[0]}
	case 3:
		res = &TreeNode{Val: nums[1]}
		res.Left = &TreeNode{Val: nums[0]}
		res.Right = &TreeNode{Val: nums[2]}
	default:
		mid := len(nums) / 2
		res = &TreeNode{Val: nums[mid]}
		res.Left = sortedArrayToBST(nums[0:mid])
		res.Right = sortedArrayToBST(nums[mid+1:])
	}
	return res
}
func main() {
	//
	nums := []int{}
	fmt.Printf("While nums is %v, the tree is \n%v\n", nums, sortedArrayToBST(nums))
}
