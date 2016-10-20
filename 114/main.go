package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	var tail *TreeNode
	if root.Left != nil {
		tail = root.Left
		for tail.Right != nil {
			tail = tail.Right
		}
		tail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}

func (t *TreeNode) String() string {
	var (
		left  string = "<nil>"
		right string = "<nil>"
	)
	if t.Left != nil {
		left = t.Left.String()
	}
	if t.Right != nil {
		right = t.Right.String()
	}
	return fmt.Sprintf("{Val: %d Left: %s Right: %s}", t.Val, left, right)
}

func main() {
	// test
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Printf("root: %v\n", root)
	flatten(root)
	fmt.Printf("flattened root: %v\n", root)
}
