package main

import "fmt"

func main() {
	fmt.Println(recoverFromPreorder("1-2--3--4-5--6--7"))
	fmt.Println(recoverFromPreorder("1-2--3---4-5--6---7"))
	fmt.Println(recoverFromPreorder("1-401--349---90--88"))
}

/*
https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/

We run a preorder depth first search on the root of a binary tree.
At each node in this traversal, we output D dashes (where D is the depth of this node), then we output the value
of this node.  (If the depth of a node is D, the depth of its immediate child is D+1.  The depth of the root node is 0.)
If a node has only one child, that child is guaranteed to be the left child.

Given the output S of this traversal, recover the tree and return its root.

Example 1:

Input: "1-2--3--4-5--6--7"
Output: [1,2,5,3,4,6,7]

Example 2:

Input: "1-2--3---4-5--6---7"
Output: [1,2,5,3,null,6,null,4,null,7]

Example 3:

Input: "1-401--349---90--88"
Output: [1,401,null,349,88,90]

Note:

The number of nodes in the original tree is between 1 and 1000.
Each node will have a value between 1 and 10^9.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	return fmt.Sprintf("%d %v %v ", tn.Val, tn.Left, tn.Right)
}

func recoverFromPreorder(S string) *TreeNode {
	root := &TreeNode{}

	if S == "" {
		return root
	}

	v, s := extractValue(&S, 0)
	root.Val = v
	recoverLayer(&S, root, 1, s)

	return root
}

func recoverLayer(s *string, parent *TreeNode, layer, start int) int {
	lr := identifyLayer(s, start)
	if lr != layer {
		return start
	}

	leftNode := &TreeNode{}
	leftNode.Val, start = extractValue(s, start+lr)
	parent.Left = leftNode

	lr = identifyLayer(s, start)
	if lr > layer {
		start = recoverLayer(s, leftNode, layer+1, start)
		lr = identifyLayer(s, start)
	}

	if lr < layer {
		return start
	}

	rightNode := &TreeNode{}
	rightNode.Val, start = extractValue(s, start+layer)
	parent.Right = rightNode

	return recoverLayer(s, rightNode, layer+1, start)
}

func extractValue(s *string, start int) (int, int) {
	value := 0
	for ; start < len(*s); start++ {
		if (*s)[start] == '-' {
			break
		}
		value = value*10 + int((*s)[start]-'0')
	}

	return value, start
}

func identifyLayer(s *string, start int) int {
	i := 0
	for i = start; i < len(*s); i++ {
		if (*s)[i] != '-' {
			break
		}
	}

	return i - start
}
