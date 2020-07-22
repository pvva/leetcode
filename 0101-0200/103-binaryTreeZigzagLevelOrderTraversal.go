package main

/*
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right,
then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	layers := [][]int{}
	cLayer := []*TreeNode{root}
	for len(cLayer) > 0 {
		nLayer := []*TreeNode{}
		cLayerVals := []int{}

		for _, node := range cLayer {
			cLayerVals = append(cLayerVals, node.Val)
			if node.Left != nil {
				nLayer = append(nLayer, node.Left)
			}
			if node.Right != nil {
				nLayer = append(nLayer, node.Right)
			}
		}

		layers = append(layers, cLayerVals)
		cLayer = nLayer
	}

	for i := 1; i < len(layers); i += 2 {
		reverseInts(layers[i])
	}

	return layers
}

func reverseInts(ints []int) {
	l := 0
	r := len(ints) - 1

	for l < r {
		ints[l], ints[r] = ints[r], ints[l]
		l++
		r--
	}
}
