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

func str(nodes ...*TreeNode) string {
	r := ""
	nextNodes := []*TreeNode{}
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == nil {
			r += "null"
		} else {
			r += fmt.Sprintf("%d", nodes[i].Val)
			nextNodes = append(nextNodes, nodes[i].Left, nodes[i].Right)
		}
		if i+1 < len(nodes) {
			r += ", "
		}
	}
	if len(nextNodes) > 0 {
		r += ", "
		r += str(nextNodes...)
	}
	return r
}

func (t *TreeNode) String() string {
	return str(t)
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
	}
	if isValidBST(root.Left) && isValidBST(root.Right) {
		if root.Left != nil {
			leftMax := root.Left
			for leftMax.Right != nil {
				leftMax = leftMax.Right
			}
			if leftMax.Val >= root.Val {
				return false
			}
		}
		if root.Right != nil {
			rightMin := root.Right
			for rightMin.Left != nil {
				rightMin = rightMin.Left
			}
			if rightMin.Val <= root.Val {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	// test
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	fmt.Printf("root: %v\nisValidBST: %t\n", root, isValidBST(root))
}
