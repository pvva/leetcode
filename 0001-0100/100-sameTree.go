package main

func main() {
	tree1 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		//Right: &TreeNode{Val: 3},
	}
	tree2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
		//Right: &TreeNode{Val: 2},
		//Left:  &TreeNode{Val: 3},
	}

	//println(isSameTree(tree1, tree2), true)
	println(isSameTree(tree1, tree2), false)
}

/*
https://leetcode.com/problems/same-tree/

Given two binary trees, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:

Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true

Example 2:

Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false

Example 3:

Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil || q == nil) || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
