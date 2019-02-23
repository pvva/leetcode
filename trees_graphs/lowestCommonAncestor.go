package main

import "fmt"

func main() {
	//p := &TreeNode{
	//	5,
	//	&TreeNode{
	//		6,
	//		nil,
	//		nil,
	//	},
	//	&TreeNode{
	//		2,
	//		&TreeNode{
	//			7,
	//			nil,
	//			nil,
	//		},
	//		&TreeNode{
	//			4,
	//			nil,
	//			nil,
	//		},
	//	},
	//}
	//q := &TreeNode{
	//	1,
	//	&TreeNode{
	//		0,
	//		nil,
	//		nil,
	//	},
	//	&TreeNode{
	//		8,
	//		nil,
	//		nil,
	//	},
	//}
	//root := &TreeNode{
	//	3,
	//	p,
	//	q,
	//}

	p := &TreeNode{
		8,
		nil,
		nil,
	}
	q := &TreeNode{
		4,
		nil,
		nil,
	}
	root := &TreeNode{
		-1,
		&TreeNode{
			0,
			&TreeNode{
				-2,
				p,
				nil,
			},
			q,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	fmt.Println(lowestCommonAncestorTrySimpler(root, p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
// The lowest common ancestor is defined between two nodes p and q as the lowest node in T that
// has both p and q as descendants (where we allow a node to be a descendant of itself)

// 1st is stupid direct approach via recursion
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val {
		return p
	}
	if root.Val == q.Val {
		return q
	}
	n, found := lca(root.Left, p, q)
	if n != nil {
		if found {
			return n
		}

		t := p
		if n == p {
			t = q
		}
		if isInBranch(root.Right, t) {
			return root
		}

		return nil
	}
	n, found = lca(root.Right, p, q)

	if n != nil && found {
		return n
	}

	return nil
}

func lca(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if root.Val == p.Val {
		return p, isInBranch(root, q)
	}
	if root.Val == q.Val {
		return q, isInBranch(root, p)
	}

	nL, foundL := lca(root.Left, p, q)
	if nL != nil {
		if foundL {
			return nL, true
		}
		t := p
		if nL == p {
			t = q
		}
		if isInBranch(root.Right, t) {
			return root, true
		}
		return nL, false
	}
	return lca(root.Right, p, q)
}

func isInBranch(b, t *TreeNode) bool {
	if b == nil {
		return false
	}

	if b.Val == t.Val {
		return true
	}

	return isInBranch(b.Left, t) || isInBranch(b.Right, t)
}

// 2nd comes as simplification of the 1st
func lowestCommonAncestorTrySimpler(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val {
		return p
	}
	if root.Val == q.Val {
		return q
	}

	l := lowestCommonAncestorTrySimpler(root.Left, p, q)
	r := lowestCommonAncestorTrySimpler(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l == nil && r == nil {
		return nil
	}

	if l != nil {
		return l
	}

	return r
}
