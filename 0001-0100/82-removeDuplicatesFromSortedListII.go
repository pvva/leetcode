package main

import "fmt"

func main() {
	list := makeList([]int{1, 2, 3, 3, 4, 4, 5})
	printList(list)
	list = deleteDuplicates(list)
	printList(list)

	list = makeList([]int{1, 1, 1, 2, 3})
	printList(list)
	list = deleteDuplicates(list)
	printList(list)

	list = makeList([]int{1, 1, 1})
	printList(list)
	list = deleteDuplicates(list)
	printList(list)
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
	p.Next = nil

	return h
}

func printList(l1 *ListNode) {
	for l1 != nil {
		fmt.Print(l1.Val)
		l1 = l1.Next
	}
	fmt.Println()
}

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct
numbers from the original list. Return the linked list sorted as well.

Example 1:


Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Example 2:

Input: head = [1,1,1,2,3]
Output: [2,3]

Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	h := &ListNode{}
	th := h

	n := head
	for n != nil {
		v := n.Val
		skip := false
		for n.Next != nil && n.Next.Val == v {
			n = n.Next
			skip = true
		}
		if skip {
			n = n.Next
			continue
		}
		if n != nil {
			h.Next = &ListNode{Val: v}
			h = h.Next
			n = n.Next
		}
	}

	return th.Next
}
