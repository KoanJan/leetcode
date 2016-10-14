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

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sum(root, []int{root.Val})
}

func sum(root *TreeNode, base []int) int {
	if root.Left == nil && root.Right == nil {
		var (
			res int = 0
			d   int = 1
		)
		for i := len(base) - 1; i >= 0; i-- {
			res += base[i] * d
			d *= 10
		}
		return res
	}

	var left, right int = 0, 0
	if root.Left != nil {
		left = sum(root.Left, append(base, root.Left.Val))
	}
	if root.Right != nil {
		right = sum(root.Right, append(base, root.Right.Val))
	}

	return left + right
}

func (t *TreeNode) String() string {
	var (
		left  string = "nil"
		right string = "nil"
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
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}

	fmt.Printf("root: %v\nsum: %d\n", root, sumNumbers(root))
}
