package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	//root := buildTree([]int{2, 1}, []int{2, 1})
	root := buildTree([]int{1, 2}, []int{2, 1})
	inorder(root)
	println()
	postorder(root)
}

/*
https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	inOrderIdxMap := make(map[int]int)
	for i, v := range inorder {
		inOrderIdxMap[v] = i
	}

	return buildTreeRec(inorder, inOrderIdxMap, 0, postorder, len(postorder)-1)
}

func buildTreeRec(inorder []int, inOrderIdxMap map[int]int, inOrderShift int, postorder []int, pIdx int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[pIdx]}
	inIdx := inOrderIdxMap[postorder[pIdx]] - inOrderShift

	root.Left = buildTreeRec(inorder[:inIdx], inOrderIdxMap, inOrderShift, postorder, pIdx-len(inorder)+inIdx)
	root.Right = buildTreeRec(inorder[inIdx+1:], inOrderIdxMap, inOrderShift+1+inIdx, postorder, pIdx-1)

	return root
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}

	inorder(node.Left)
	print(node.Val, " ")
	inorder(node.Right)
}

func postorder(node *TreeNode) {
	if node == nil {
		return
	}

	postorder(node.Left)
	postorder(node.Right)
	print(node.Val, " ")
}
