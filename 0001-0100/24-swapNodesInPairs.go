package main

import "fmt"

func main() {
	l1 := makeList([]int{1, 2, 3, 4})
	printList(l1)
	printList(swapPairs(l1))
	l1 = makeList([]int{})
	printList(l1)
	printList(swapPairs(l1))
	l1 = makeList([]int{1})
	printList(l1)
	printList(swapPairs(l1))
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
https://leetcode.com/problems/swap-nodes-in-pairs/

Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without
modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Example 1:

Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:

Input: head = []
Output: []

Example 3:

Input: head = [1]
Output: [1]
*/

func swapPairs(head *ListNode) *ListNode {
	t := &ListNode{}
	th := t

	h := head
	for h != nil && h.Next != nil {
		nh := h.Next.Next
		t.Next = h.Next
		t.Next.Next = h
		t = t.Next.Next
		t.Next = nil
		h = nh
	}
	if h != nil {
		t.Next = h
	}

	return th.Next
}
