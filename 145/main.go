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

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	left := postorderTraversal(root.Left)
	for i := 0; i < len(left); i++ {
		res = append(res, left[i])
	}
	right := postorderTraversal(root.Right)
	for i := 0; i < len(right); i++ {
		res = append(res, right[i])
	}
	res = append(res, root.Val)
	return res
}

func main() {
	// test
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Printf("postorderTraversal of tree %v is %v\n", root, postorderTraversal(root))
}
