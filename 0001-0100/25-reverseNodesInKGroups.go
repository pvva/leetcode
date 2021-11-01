package main

import "fmt"

func main() {
	list := makeList([]int{1, 2, 3, 4, 5})
	printList(list)
	list = reverseKGroup(list, 3)
	printList(list)

	list = makeList([]int{1, 2, 3, 4, 5})
	printList(list)
	list = reverseKGroup(list, 2)
	printList(list)

	list = makeList([]int{1, 2, 3, 4, 5})
	printList(list)
	list = reverseKGroup(list, 1)
	printList(list)

	list = makeList([]int{1})
	printList(list)
	list = reverseKGroup(list, 1)
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
https://leetcode.com/problems/reverse-nodes-in-k-group/

Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is
not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]

Example 2:

Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]

Example 3:

Input: head = [1,2,3,4,5], k = 1
Output: [1,2,3,4,5]

Example 4:

Input: head = [1], k = 1
Output: [1]

Constraints:

The number of nodes in the list is in the range sz.
1 <= sz <= 5000
0 <= Node.val <= 1000
1 <= k <= sz


Follow-up: Can you solve the problem in O(1) extra memory space?
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	var h *ListNode
	var sn *ListNode
	l := head
	for l != nil {
		st := l
		c := k
		var t *ListNode
		for c > 0 && l != nil {
			ln := l
			l = l.Next

			ln.Next = t
			t = ln
			c--
		}
		if c == 0 {
			if h == nil {
				h = t
			} else {
				sn.Next = t
			}
			sn = st
		} else {
			l = nil
			for t != nil {
				ln := t
				t = t.Next

				ln.Next = l
				l = ln
			}
			sn.Next = l
			break
		}
	}

	return h
}
