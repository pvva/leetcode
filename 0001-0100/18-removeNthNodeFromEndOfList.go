package main

import "fmt"

func main() {
	l := makeList([]int{1, 2, 3, 4, 5})
	printList(l)
	printList(removeNthFromEnd2(l, 2))

	l = makeList([]int{1})
	printList(l)
	printList(removeNthFromEnd2(l, 1))

	l = makeList([]int{1, 2})
	printList(l)
	printList(removeNthFromEnd2(l, 1))
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
https://leetcode.com/problems/remove-nth-node-from-end-of-list/

Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:

Input: head = [1], n = 1
Output: []

Example 3:

Input: head = [1,2], n = 1
Output: [1]
*/

func removeFromEnd(lp **ListNode, n int) int {
	l := *lp
	c := 0
	if l.Next != nil {
		c = removeFromEnd(&l.Next, n) + 1
	}
	if c == n-1 {
		if l.Next != nil {
			*lp = (*lp).Next
		} else {
			*lp = nil
		}
	}
	return c
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	removeFromEnd(&head, n)
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	p1 := head
	p2 := &head

	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = &(*p2).Next
	}
	if (*p2).Next == nil {
		*p2 = nil
	} else {
		*p2 = (*p2).Next
	}

	return head
}
