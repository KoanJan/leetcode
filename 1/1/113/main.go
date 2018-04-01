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

func pathSum(root *TreeNode, sum int) [][]int {
	return ps(root, sum, []int{})
}

func ps(root *TreeNode, sum int, prePath []int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		res = append(res, append(cpSlice(prePath), sum))
	} else if root.Left != nil || root.Right != nil {
		if root.Left != nil {
			leftSubRes := ps(root.Left, sum-root.Val, append(cpSlice(prePath), root.Val))
			for i := 0; i < len(leftSubRes); i++ {
				res = append(res, cpSlice(leftSubRes[i]))
			}
		}
		if root.Right != nil {
			rightSubRes := ps(root.Right, sum-root.Val, append(cpSlice(prePath), root.Val))
			for i := 0; i < len(rightSubRes); i++ {
				res = append(res, cpSlice(rightSubRes[i]))
			}
		}
	}
	return res
}

func cpSlice(a []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i]
	}
	return res
}

func main() {
	// test
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}
	sum := 22

	fmt.Printf("root: %v\nsum: %d\n\npaths match sum: %v\n", root, sum, pathSum(root, sum))
}
