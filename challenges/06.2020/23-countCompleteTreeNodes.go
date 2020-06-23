package main

/*
Given a complete binary tree, count the number of nodes.

Note:

Definition of a complete binary tree from Wikipedia:
In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level
are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Example:

Input:
    1
   / \
  2   3
 / \  /
4  5 6

Output: 6
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// idea is to measure the height from the left and from the right subtree
// if heights are equal, then number of nodes is 2^h - 1 (it is the sum of the row 2^0 + 2^1 + 2^2 + .. 2^(h-1))
// if heights are not equal just recurse to the next level of the tree and count there
// due to recursion + height calculation time complexity is O(log^2(n)), where n is the number of nodes
func countNodes(root *TreeNode) int {
	lh := treeHeight(root, true)
	rh := treeHeight(root, false)

	if lh == rh {
		return 1<<lh - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func treeHeight(node *TreeNode, left bool) int {
	h := 0
	n := node
	for n != nil {
		if left {
			n = n.Left
		} else {
			n = n.Right
		}
		h++
	}

	return h
}
