package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{
		4,
		&ListNode{
			2,
			&ListNode{
				1,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}

	printList(list)
	printList(sortList(list))
}

func printList(head *ListNode) {
	d := ""
	p := head
	for p != nil {
		print(d, p.Val)
		d = "->"
		p = p.Next
	}
	println()
}

/*
https://leetcode.com/problems/sort-list/

Sort a linked list in O(n log n) time using constant space complexity.

Example 1:

Input: 4->2->1->3
Output: 1->2->3->4
Example 2:

Input: -1->5->3->4->0
Output: -1->0->3->4->5
*/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	h := head
	e := head
	for e != nil && e.Next != nil {
		p = h
		h = h.Next
		e = e.Next.Next
	}
	p.Next = nil

	return merge(sortList(head), sortList(h))
}

func merge(l1, l2 *ListNode) *ListNode {
	var h, n *ListNode
	p1 := l1
	p2 := l2

	for p1 != nil || p2 != nil {
		var next *ListNode
		if p1 == nil {
			next = p2
			p2 = p2.Next
		} else if p2 == nil {
			next = p1
			p1 = p1.Next
		} else {
			if p1.Val <= p2.Val {
				next = p1
				p1 = p1.Next
			} else {
				next = p2
				p2 = p2.Next
			}
		}

		if n == nil {
			n = next
			h = next
		} else {
			n.Next = next
			n = n.Next
		}
	}

	return h
}
