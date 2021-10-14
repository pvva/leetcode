package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 3},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{Val: 4},
		},
	}
	printList(l1)
	printList(l2)
	printList(addTwoNumbers(l1, l2))

	l1 = &ListNode{Val: 0}
	l2 = &ListNode{Val: 0}
	printList(l1)
	printList(l2)
	printList(addTwoNumbers(l1, l2))

	l1 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}
	l2 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	printList(l1)
	printList(l2)
	printList(addTwoNumbers(l1, l2))
}

func printList(l1 *ListNode) {
	for l1 != nil {
		fmt.Print(l1.Val)
		l1 = l1.Next
	}
	fmt.Println()
}

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	n := res
	c := 0
	for l1 != nil || l2 != nil {
		v := c
		c = 0
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if v >= 10 {
			c = 1
			v -= 10
		}
		if n.Next == nil {
			n.Next = &ListNode{Val: v}
			n = n.Next
		}
	}
	if c > 0 {
		n.Next = &ListNode{Val: c}
	}

	return res.Next
}
