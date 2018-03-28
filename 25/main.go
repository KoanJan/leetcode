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

func reverseKGroup(head *ListNode, k int) *ListNode {
	root := &ListNode{Next: head}
	cur := root
	var gHead, gTail *ListNode
	for cur.Next != nil {
		// find group head and tail
		gHead = cur
		for i := 0; i < k; i++ {
			if cur.Next == nil {
				return root.Next
			}
			cur = cur.Next
		}
		gTail = cur
		cGTail := gTail
		// reverse
		for i := 0; i < k-1; i++ {
			cGTail = gTail
			rem := cGTail.Next
			cGTail.Next = gHead.Next
			cGTail = cGTail.Next
			gHead.Next = gHead.Next.Next
			cGTail.Next = rem
		}
		// fix current
		for i := 0; i < k-1; i++ {
			cur = cur.Next
		}
	}
	return root.Next
}

func newNode(val ...int) *ListNode {
	node := new(ListNode)
	head := node
	for _, v := range val {
		n := &ListNode{v, nil}
		node.Next = n
		node = node.Next
	}
	return head.Next
}

func printListNode(node *ListNode) {
	c := node
	for c != nil {
		fmt.Printf("%d ", c.Val)
		c = c.Next
	}
	fmt.Println("")
}

func main() {
	printListNode(reverseKGroup(newNode(1, 2, 3, 4, 5, 6, 7), 3))
}
