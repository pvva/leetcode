package main

/*
 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL

*/

func main() {
	n12 := &Node{Val: 12}
	n11 := &Node{Val: 11, Next: n12}
	n12.Prev = n11

	n10 := &Node{Val: 10}
	n9 := &Node{Val: 9, Next: n10}
	n10.Prev = n9
	n8 := &Node{Val: 8, Next: n9}
	n9.Prev = n8
	n7 := &Node{Val: 7, Next: n8}
	n8.Prev = n7
	n8.Child = n11

	n6 := &Node{Val: 6}
	n5 := &Node{Val: 5, Next: n6}
	n6.Prev = n5
	n4 := &Node{Val: 4, Next: n5}
	n5.Prev = n4
	n3 := &Node{Val: 3, Next: n4}
	n4.Prev = n3
	n2 := &Node{Val: 2, Next: n3}
	n3.Prev = n2
	n1 := &Node{Val: 1, Next: n2}
	n3.Child = n7

	//flatten(n1)
	flattenNonRec(n1)
	printList(n1)
}

func printList(n *Node) {
	for n != nil {
		print(n.Val, " ")
		n = n.Next
	}
	println()
}

/*
https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/

You are given a doubly linked list which in addition to the next and previous pointers, it could have a child pointer,
which may or may not point to a separate doubly linked list. These child lists may have one or more children of their
own, and so on, to produce a multilevel data structure, as shown in the example below.

Flatten the list so that all the nodes appear in a single-level, doubly linked list. You are given the head of the first
level of the list.

Example 1:

Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
Output: [1,2,3,7,8,11,12,9,10,4,5,6]

Explanation:
The input multilevel linked list is as follows:

 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL

Example 2:

Input: head = [1,2,null,3]
Output: [1,3,2]

Explanation:
The input multilevel linked list is as follows:

  1---2---NULL
  |
  3---NULL

Example 3:

Input: head = []
Output: []
*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	n, _ := flattenSub(root)

	return n
}

func flattenSub(node *Node) (*Node, *Node) {
	n := node
	p := node

	for n != nil {
		p = n
		if n.Child != nil {
			subF, subE := flattenSub(n.Child)
			n.Child = nil
			subE.Next = n.Next
			if n.Next != nil {
				n.Next.Prev = subE
			}
			n.Next = subF
			subF.Prev = n
			n = subE
		} else {
			n = n.Next
		}
	}

	return node, p
}

func flattenNonRec(root *Node) *Node {
	n := root
	var p *Node

	for n != nil {
		n.Prev = p
		if n.Child != nil {
			le := n.Child
			for le.Next != nil {
				le = le.Next
			}
			le.Next = n.Next

			n.Next = n.Child
			n.Child.Prev = n

			n.Child = nil
		}
		p = n
		n = n.Next
	}

	return root
}
