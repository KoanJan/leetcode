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

type linkedList struct {
	Val  int
	Next *linkedList
}

func newLinkedList(n int) *linkedList {
	l := &linkedList{Val: n}
	return l
}

func (l *linkedList) Push(n ...int) {
	if len(n) > 0 {
		tail := l
		for tail.Next != nil {
			tail = tail.Next
		}
		for _, e := range n {
			tail.Next = new(linkedList)
			tail.Next.Val = e
			tail = tail.Next
		}
	}
}

func (l *linkedList) Unshift(n ...int) *linkedList {
	if len(n) > 0 {
		newList := newLinkedList(n[0])
		newList.Push(n...)
		newList = newList.Next
		tail := newList
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = l
		return newList
	}
	return l
}

func (l *linkedList) Array() []int {
	var (
		res = []int{}
		cur = l
	)
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func (l *linkedList) String() string {
	var (
		res string = ""
		c          = l
	)
	for c != nil {
		res = fmt.Sprintf("%s%d - ", res, c.Val)
		c = c.Next
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		res = newLinkedList(root.Val)
	)
	if root.Left != nil {
		res = res.Unshift(inorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res.Push(inorderTraversal(root.Right)...)
	}
	return res.Array()

}

func main() {

	// test
	tree := &TreeNode{Val: 1}
	tree.Left = nil
	right := &TreeNode{Val: 2}
	right.Left = &TreeNode{Val: 3}
	tree.Right = right

	fmt.Printf("inorderTraversal(tree) = %v\n", inorderTraversal(tree))
}
