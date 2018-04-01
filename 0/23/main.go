package main

import (
	"container/heap"
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

func mergeKLists(lists []*ListNode) *ListNode {
	lists = filtNil(lists)
	node := new(ListNode)
	head := node
	q := newPriorityQueue()
	for _, list := range lists {
		q.Enqueue(list)
	}
	for !q.IsEmpty() {
		list := q.Dequeue()
		node.Next = list
		node = node.Next
		if node.Next != nil {
			q.Enqueue(node.Next)
			node.Next = nil
		}
	}
	return head.Next

}

func filtNil(lists []*ListNode) []*ListNode {
	if len(lists) == 0 {
		return lists
	}
	empty := 0
	for i := range lists {
		if lists[i] == nil {
			lists[i], lists[empty] = lists[empty], lists[i]
			empty++
		}
	}
	return lists[empty:]
}

type PriorityQueue struct {
	data *priorityQueue
}

func newPriorityQueue() *PriorityQueue {
	queue := new(PriorityQueue)
	queue.data = new(priorityQueue)
	*queue.data = make([]*ListNode, 0)
	return queue
}

func (queue *PriorityQueue) Enqueue(node *ListNode) {
	heap.Push(queue.data, node)
}

func (queue *PriorityQueue) Dequeue() *ListNode {
	return heap.Pop(queue.data).(*ListNode)
}

func (queue *PriorityQueue) IsEmpty() bool {
	return queue.data.Len() == 0
}

type priorityQueue []*ListNode

func (queue priorityQueue) Len() int {
	return len(queue)
}

func (queue priorityQueue) Less(i, j int) bool {
	return queue[i].Val < queue[j].Val
}

func (queue priorityQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
}

func (queue *priorityQueue) Push(x interface{}) {
	*queue = append(*queue, x.(*ListNode))
}

func (queue *priorityQueue) Pop() interface{} {
	n := len(*queue)
	x := (*queue)[n-1]
	*queue = (*queue)[0 : n-1]
	return x
}

func printListNode(node *ListNode) {
	c := node
	for c != nil {
		fmt.Printf("%d ", c.Val)
		c = c.Next
	}
}

func main() {
	lists := []*ListNode{
		newNode(1, 3, 5),
		newNode(0, 4),
	}
	m := mergeKLists(lists)
	printListNode(m)
}
