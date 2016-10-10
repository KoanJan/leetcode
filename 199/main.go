package main

import (
	"fmt"
)

// TreeNode. Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		input  = []*TreeNode{root}
		output = []int{}
	)
	for len(input) > 0 {
		output = append(output, rightSideNode(input...))
		newInput := []*TreeNode{}
		for i := 0; i < len(input); i++ {
			if input[i] != nil {
				if input[i].Left != nil {
					newInput = append(newInput, input[i].Left)
				}
				if input[i].Right != nil {
					newInput = append(newInput, input[i].Right)
				}
			}
		}
		input = newInput
	}
	return output
}

func rightSideNode(nodes ...*TreeNode) int {
	if len(nodes) == 0 {
		panic("Invalid input")
	}
	return nodes[len(nodes)-1].Val
}

func main() {

	// test
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	fmt.Printf("rightSideView of tree %v is %v\n", root, rightSideView(root))
}
