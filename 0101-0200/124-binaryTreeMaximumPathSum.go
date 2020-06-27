package main

import "math"

func main() {
	//r := &TreeNode{
	//	1,
	//	&TreeNode{2, nil, nil},
	//	&TreeNode{3, nil, nil},
	//}

	//r := &TreeNode{
	//	-10,
	//	&TreeNode{9, nil, nil},
	//	&TreeNode{
	//		20,
	//		&TreeNode{15, nil, nil},
	//		&TreeNode{7, nil, nil},
	//	},
	//}

	//r := &TreeNode{
	//	1,
	//	&TreeNode{-2,
	//		&TreeNode{
	//			1,
	//			&TreeNode{-1, nil, nil},
	//			nil},
	//		&TreeNode{3, nil, nil},
	//	},
	//	&TreeNode{
	//		-3,
	//		&TreeNode{-2, nil, nil},
	//		nil,
	//	},
	//}

	r := &TreeNode{
		-1,
		nil,
		&TreeNode{
			9,
			&TreeNode{-6, nil, nil},
			&TreeNode{
				3,
				nil,
				&TreeNode{-2, nil, nil},
			},
		},
	}

	println(maxPathSum(r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
https://leetcode.com/problems/binary-tree-maximum-path-sum/

Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along
the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6
Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
*/

func maxPathSum(root *TreeNode) int {
	return max(mps(root))
}

// returns (max path value that can be continued, max path value that ends on this node or close to this node)
func mps(node *TreeNode) (m, cm int) {
	m = math.MinInt32
	cm = math.MinInt32

	if node == nil {
		return
	}
	lm, lcm := mps(node.Left)
	rm, rcm := mps(node.Right)

	m = max(node.Val, max(lm+node.Val, rm+node.Val))
	cm = max(lm, max(rm, max(node.Val, max(lcm, max(rcm, lm+rm+node.Val)))))

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
