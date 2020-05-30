package main

import (
	"strconv"
	"strings"
)

func main() {
	l1 := &ListNode{
		1,
		&ListNode{
			2,
			nil,
		},
	}
	l2 := &ListNode{
		1,
		&ListNode{
			2,
			nil,
		},
	}
	l3 := &ListNode{
		1,
		&ListNode{
			2,
			nil,
		},
	}
	l4 := &ListNode{
		0,
		&ListNode{
			1,
			&ListNode{
				2,
				nil,
			},
		},
	}

	println(l1.String(), "|", rotateRight(l1, 1).String())
	println(l2.String(), "|", rotateRight(l2, 2).String())
	println(l3.String(), "|", rotateRight(l3, 3).String())
	println(l4.String(), "|", rotateRight(l4, 4).String())
}

/*
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL

Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL

Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	sb := strings.Builder{}
	d := ""
	p := ln
	for p != nil {
		sb.WriteString(d)
		sb.WriteString(strconv.Itoa(p.Val))
		d = " -> "
		p = p.Next
	}

	return sb.String()
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	size := 1

	cPtr := head
	kPtr := head

	for k > 0 && cPtr.Next != nil {
		size++
		k--
		cPtr = cPtr.Next
	}

	if k > 0 {
		if size == 1 {
			return head
		}
		return rotateRight(head, (k-1)%size)
	}
	kPtr = head
	for cPtr.Next != nil {
		cPtr = cPtr.Next
		kPtr = kPtr.Next
	}

	cPtr.Next = head
	ret := kPtr.Next
	kPtr.Next = nil

	return ret
}
