package main

func main() {
	list := &ListNode{Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 12,
						Next: &ListNode{
							Val: 16,
						},
					},
				},
			},
		},
	}
	//printList(list)
	//rev := reverse(list)
	//printList(rev)

	original := restoreList(list)
	printList(original)
	println("1 8 2 9 16 12")
}

/*
You are given a singly-linked list that contains N integers. A subpart of the list is a contiguous set of even elements,
bordered either by the end of the list or an odd element. For example, if the list is [1, 2, 8, 9, 12, 16], the subparts
of the list are [2, 8] and [12, 16].
Then, for each subpart, the order of the elements is reversed. In the example, this would result in the new list,
[1, 8, 2, 9, 16, 12].
The goal of this question is: given a resulting list, determine the original order of the elements.

Constraints
1 <= N <= 1000, where N is the size of the list
1 <= Li <= 10^9, where Li is the ith element of the list

Example
Input:
N = 6
list = [1, 2, 8, 9, 12, 16]
Output:
[1, 8, 2, 9, 16, 12]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func restoreList(node *ListNode) *ListNode {
	// the idea is to have 2 pointers for reversed list and do in place reverse and 2 pointers for resulting list,
	// one head and one current
	// O(n)
	var cHead, cLast, head, current *ListNode
	for node != nil {
		next := node.Next
		if node.Val&1 == 0 {
			// even
			if cHead == nil {
				cHead = node
				cLast = node
				cLast.Next = nil
			} else {
				t := cHead
				cHead = node
				cHead.Next = t
			}
		} else {
			// odd
			if head == nil {
				if cHead == nil {
					head = node
				} else {
					head = cHead
				}
			}

			if cHead != nil {
				if current != nil {
					current.Next = cHead
				}
				cHead = nil
				cLast.Next = node
				current = cLast.Next
			} else {
				if current != nil {
					current.Next = node
					current = current.Next
				} else {
					current = node
				}
			}
		}

		node = next
	}
	if cHead != nil && current != nil {
		current.Next = cHead
	}
	if head == nil {
		return cHead
	}

	return head
}

func reverse(node *ListNode) *ListNode {
	var cList *ListNode

	for node != nil {
		t := node.Next
		node.Next = cList
		cList = node
		node = t
	}

	return cList
}

func printList(node *ListNode) {
	n := node
	for n != nil {
		print(n.Val, " ")
		n = n.Next
	}
	println()
}
