package main

import "fmt"

func main() {
	list := makeList([]int{1, 4, 3, 2, 5, 2})
	printList(list)
	list = partition(list, 3)
	printList(list)

	list = makeList([]int{2, 1})
	printList(list)
	list = partition(list, 2)
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
https://leetcode.com/problems/partition-list/

Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater
than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example 1:


Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]

Example 2:

Input: head = [2,1], x = 2
Output: [1,2]

Constraints:

The number of nodes in the list is in the range [0, 200].
-100 <= Node.val <= 100
-200 <= x <= 200
*/

func partition(head *ListNode, x int) *ListNode {
	newHeadStore := &ListNode{}
	currentLessNode := newHeadStore
	currentNode := head
	currentGreaterNode := &ListNode{}
	greaterHeadStore := currentGreaterNode
	for currentNode != nil {
		nextNode := currentNode.Next
		if currentNode.Val < x {
			currentLessNode.Next = currentNode
			currentLessNode = currentLessNode.Next
			currentLessNode.Next = nil
		} else {
			currentGreaterNode.Next = currentNode
			currentGreaterNode = currentGreaterNode.Next
			currentGreaterNode.Next = nil
		}
		currentNode = nextNode
	}
	currentLessNode.Next = greaterHeadStore.Next

	return newHeadStore.Next
}
