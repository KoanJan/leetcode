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

// recursive
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return combineSlice([]int{root.Val}, preorderTraversal(root.Left), preorderTraversal(root.Right))
}

func combineSlice(slices ...[]int) []int {
	res := []int{}
	for _, slice := range slices {
		for _, elem := range slice {
			res = append(res, elem)
		}
	}
	return res
}

func main() {

	// test
	tree := &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 2}
	tree.Right.Left = &TreeNode{Val: 3}

	fmt.Printf("preorderTraversal result: %v", preorderTraversal(tree))
}
