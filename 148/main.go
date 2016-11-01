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

func newListNode(nums []int) *ListNode {
	var (
		res = new(ListNode)
		cur = res
	)
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return res.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var (
		head1 *ListNode = head
		head2 *ListNode = getMid(head)
	)
	head1 = sortList(head1)
	head2 = sortList(head2)
	return merge(head1, head2)
}

func merge(head1, head2 *ListNode) *ListNode {
	var (
		head *ListNode = new(ListNode)
		tail *ListNode = head
	)
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			tail.Next = &ListNode{Val: head1.Val}
			head1 = head1.Next
		} else {
			tail.Next = &ListNode{Val: head2.Val}
			head2 = head2.Next
		}
		tail = tail.Next
		tail.Next = nil
	}
	if head1 != nil {
		tail.Next = head1
	}
	if head2 != nil {
		tail.Next = head2
	}
	return head.Next
}

func getMid(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var (
		fast, slow, prev *ListNode = head.Next, head.Next, head
	)
	for {
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}
		prev = slow
		slow = slow.Next
	}
	prev.Next = nil
	return slow
}

func main() {
	// test
	head := newListNode([]int{5, 1, 6, 3, 8, 7, 4})
	fmt.Printf("head: %v\n", head)
	head = sortList(head)
	fmt.Printf("sorted head: %v\n", head)
}
