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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return zigzag(true, root)
}

func zigzag(isLeft bool, nodes ...*TreeNode) [][]int {
	res := [][]int{}
	if len(nodes) > 0 {
		var (
			thisLevel = []int{}
			nextNodes = []*TreeNode{}
		)
		if isLeft {
			for i := 0; i < len(nodes); i++ {
				if nodes[i] != nil {
					thisLevel = append(thisLevel, nodes[i].Val)
				}
			}
		} else {
			for i := len(nodes) - 1; i >= 0; i-- {
				if nodes[i] != nil {
					thisLevel = append(thisLevel, nodes[i].Val)
				}
			}
		}
		for i := 0; i < len(nodes); i++ {
			if nodes[i] != nil {
				if nodes[i].Left != nil {
					nextNodes = append(nextNodes, nodes[i].Left)
				}
				if nodes[i].Right != nil {
					nextNodes = append(nextNodes, nodes[i].Right)
				}
			}
		}
		res = append(res, thisLevel)
		res = append(res, zigzag(!isLeft, nextNodes...)...)
	}
	return res
}

func main() {
	// test
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Printf("root: %v, zigzag: %v\n", root, zigzagLevelOrder(root))
}
