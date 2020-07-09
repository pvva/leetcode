package main

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 9},
		},
	}

	println(widthOfBinaryTree2(root), 4)

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 3},
		},
	}

	println(widthOfBinaryTree2(root), 2)

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 2},
	}

	println(widthOfBinaryTree2(root), 2)

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 6},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   9,
				Right: &TreeNode{Val: 7},
			},
		},
	}

	println(widthOfBinaryTree2(root), 8)
}

/*
https://leetcode.com/problems/maximum-width-of-binary-tree/

Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum
width among all levels. The binary tree has the same structure as a full binary tree, but some nodes are null.

The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes
in the level, where the null nodes between the end-nodes are also counted into the length calculation.

Example 1:

Input:

           1
         /   \
        3     2
       / \     \
      5   3     9

Output: 4
Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).

Example 2:

Input:

          1
         /
        3
       / \
      5   3

Output: 2
Explanation: The maximum width existing in the third level with the length 2 (5,3).

Example 3:

Input:

          1
         / \
        3   2
       /
      5

Output: 2
Explanation: The maximum width existing in the second level with the length 2 (3,2).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	cLayer := []*TreeNode{root}
	w := 0

	for len(cLayer) > 0 {
		nLayer := []*TreeNode{}
		l := len(cLayer) - 1
		for l >= 0 && cLayer[l] == nil {
			l--
		}
		s := 0
		for s <= l && cLayer[s] == nil {
			s++
		}
		nw := l - s + 1
		if nw > w {
			w = nw
		}

		for i := s; i <= l; i++ {
			if cLayer[i] == nil {
				nLayer = append(nLayer, nil)
				nLayer = append(nLayer, nil)
			} else {
				nLayer = append(nLayer, cLayer[i].Left)
				nLayer = append(nLayer, cLayer[i].Right)
			}
		}
		cLayer = nLayer
	}

	return w
}

func widthOfBinaryTree2(root *TreeNode) int {
	cLayer := []*TreeNode{{Val: 0, Left: root.Left, Right: root.Right}}
	w := 0

	for len(cLayer) > 0 {
		nLayer := []*TreeNode{}
		l := cLayer[len(cLayer)-1].Val - cLayer[0].Val + 1
		if l > w {
			w = l
		}

		for _, n := range cLayer {
			if n.Left != nil {
				nLayer = append(nLayer, &TreeNode{Val: n.Val * 2, Left: n.Left.Left, Right: n.Left.Right})
			}
			if n.Right != nil {
				nLayer = append(nLayer, &TreeNode{Val: n.Val*2 + 1, Left: n.Right.Left, Right: n.Right.Right})
			}
		}

		cLayer = nLayer
	}

	return w
}
