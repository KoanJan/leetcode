package main

import (
	"fmt"
)

// TreeNode: Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		next   = rob(root.Left) + rob(root.Right)
		profit = root.Val
	)
	if root.Left != nil {
		profit += (rob(root.Left.Left) + rob(root.Left.Right))
	}
	if root.Right != nil {
		profit += (rob(root.Right.Left) + rob(root.Right.Right))
	}
	if next > profit {
		return next
	}
	return profit
}

func main() {

	// test
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 1}

	fmt.Printf("rob profit: %d\n", rob(root))
}
