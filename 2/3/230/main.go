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

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		panic("not found")
	}
	leftSize := calcTreeSize(root.Left)
	if k == leftSize+1 {
		return root.Val
	} else if k > leftSize+1 {
		return kthSmallest(root.Right, k-leftSize-1)
	} else {
		return kthSmallest(root.Left, k)
	}
}

func calcTreeSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + calcTreeSize(root.Left) + calcTreeSize(root.Right)
}

func main() {

	// test
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	k := 3

	fmt.Printf("kthSmallest(root, %d) = %d\n", k, kthSmallest(root, k))
}
