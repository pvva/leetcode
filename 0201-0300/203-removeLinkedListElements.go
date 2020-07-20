package main

/*
https://leetcode.com/problems/remove-linked-list-elements/

Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var retHead, retCurr *ListNode

	for head != nil {
		if head.Val != val {
			head = head.Next
			continue
		}
		if retHead == nil {
			retHead = head
		}
		if retCurr != nil {
			retCurr.Next = head
		}
		retCurr = head

		head = head.Next
		retCurr.Next = nil
	}

	return retHead
}
