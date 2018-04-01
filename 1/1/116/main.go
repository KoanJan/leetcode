package main

import (
	"fmt"
)

// TreeLinkNode. Definition for binary tree with next pointer.
type TreeLinkNode struct {
	Val               int
	Left, Right, Next *TreeLinkNode
}

func connect(root *TreeLinkNode) {
	c(root)
}

func c(nodes ...*TreeLinkNode) {
	if len(nodes) == 0 {
		return
	}
	nodes = append(nodes, nil)
	newNodes := []*TreeLinkNode{}
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i] == nil {
			continue
		}
		nodes[i].Next = nodes[i+1]
		newNodes = append(newNodes, nodes[i].Left)
		newNodes = append(newNodes, nodes[i].Right)
	}
	if len(newNodes) != 0 {
		c(newNodes...)
	}
}

func (t *TreeLinkNode) String() string {
	return fmt.Sprintf("{Val: %d, Left: %v, Right: %v, Next: %v}", t.Val, t.Left, t.Right, t.Next)
}

func main() {

	// test
	root := &TreeLinkNode{Val: 1}
	root.Left = &TreeLinkNode{Val: 2}
	root.Right = &TreeLinkNode{Val: 3}
	root.Left.Left = &TreeLinkNode{Val: 4}
	root.Right.Left = &TreeLinkNode{Val: 5}

	fmt.Printf("tree:\n%v\n", root)
	connect(root)
	fmt.Printf("\n")
	fmt.Printf("connected:\n%v\n", root)
}
