package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
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

	r := len(layers) - 1
	l := 0
	for l < r {
		layers[l], layers[r] = layers[r], layers[l]
		r--
		l++
	}

	return layers
}
