package main

import (
	"container/heap"
	"fmt"
)

func main() {
	printList(mergeKLists([]*ListNode{
		makeList([]int{1, 4, 5}),
		makeList([]int{1, 3, 4}),
		makeList([]int{2, 6}),
	}))
	printList(mergeKLists([]*ListNode{}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(values []int) *ListNode {
	h := &ListNode{}

	c := h
	var p *ListNode
	for _, v := range values {
		c.Val = v
		c.Next = &ListNode{}
		p = c
		c = c.Next
	}
	if p != nil {
		p.Next = nil
	}

	return h
}

func printList(l1 *ListNode) {
	for l1 != nil {
		fmt.Print(l1.Val, " ")
		l1 = l1.Next
	}
	fmt.Println()
}

/*
https://leetcode.com/problems/merge-k-sorted-lists/

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Example 2:

Input: lists = []
Output: []

Example 3:

Input: lists = [[]]
Output: []

Constraints:

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] is sorted in ascending order.
The sum of lists[i].length won't exceed 10^4.
*/

type ListHeap []*ListNode

func (h ListHeap) Len() int {
	return len(h)
}
func (h ListHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h ListHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *ListHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}
func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}
	head := res
	lheap := &ListHeap{}
	for _, l := range lists {
		if l != nil {
			heap.Push(lheap, l)
		}
	}

	for lheap.Len() > 0 {
		l := heap.Pop(lheap).(*ListNode)
		res.Next = &ListNode{Val: l.Val}
		res = res.Next
		l = l.Next
		if l != nil {
			heap.Push(lheap, l)
		}
	}

	return head.Next
}
