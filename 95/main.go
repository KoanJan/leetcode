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

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{}
	}
	ran := make([]int, n)
	for i := 0; i < n; i++ {
		ran[i] = i + 1
	}
	return genTrees(ran)
}

func genTrees(ran []int) []*TreeNode {
	res := []*TreeNode{}
	if len(ran) == 0 {
		res = append(res, nil)
		return res
	}
	if len(ran) == 1 {
		res = append(res, &TreeNode{Val: ran[0]})
		return res
	}
	for i := 0; i < len(ran); i++ {
		left := genTrees(ran[0:i])
		right := genTrees(ran[i+1:])
		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				parent := &TreeNode{Val: ran[i]}
				parent.Left = left[j]
				parent.Right = right[k]
				res = append(res, parent)
			}
		}
	}
	return res
}

func main() {
	// test
	var (
		n           int         = 3
		trees       []*TreeNode = generateTrees(n)
		treesString string      = "\n"
	)
	for i := 0; i < len(trees); i++ {
		treesString = fmt.Sprintf("%s[%v]\n", treesString, trees[i])
	}

	fmt.Printf("While n = %d, trees: %s\n", n, treesString)
}
