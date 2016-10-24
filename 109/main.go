package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "<nil>"
	}
	var (
		res string
		cur *ListNode = l
	)
	for cur != nil {
		res += fmt.Sprintf("%d", cur.Val)
		if cur.Next != nil {
			res += " -> "
		}
		cur = cur.Next
	}
	return res
}

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

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var (
		length int       = 1
		cur    *ListNode = head
		res    *TreeNode
	)
	for cur.Next != nil {
		length++
		cur = cur.Next
	}
	cur = head

	if length == 1 {
		res = &TreeNode{Val: head.Val}
	} else if length == 2 {
		res = &TreeNode{Val: head.Next.Val}
		res.Left = &TreeNode{Val: head.Val}
	} else if length == 3 {
		res = &TreeNode{Val: head.Next.Val}
		res.Left = &TreeNode{Val: head.Val}
		res.Right = &TreeNode{Val: head.Next.Next.Val}
	} else {
		leftTail := cur
		for i := 1; i < length/2; i++ {
			leftTail = leftTail.Next
		}
		mid := leftTail.Next
		leftTail.Next = nil
		res = &TreeNode{Val: mid.Val}
		res.Left = sortedListToBST(cur)
		res.Right = sortedListToBST(mid.Next)
	}
	return res
}

func main() {
	// test
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	fmt.Printf("head: \n%v\n", head)
	bst := sortedListToBST(head)
	fmt.Printf("bst: \n%v\n", bst)
}
