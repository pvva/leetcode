package main

import "math"

func main() {
	root := &Node{
		Val: 8,
		Left: &Node{
			Val:   3,
			Left:  &Node{Val: 1},
			Right: &Node{Val: 6, Left: &Node{Val: 4}, Right: &Node{Val: 7}},
		},
		Right: &Node{
			Val:   10,
			Right: &Node{Val: 14, Left: &Node{Val: 13}},
		},
	}

	println(calculateMaxAbsDiff(root), 7)
}

/*
Given binary tree calculate maximum absolute difference between two nodes one of which is the ancestor of another.
*/

type Node struct {
	Val         int
	Left, Right *Node
}

func calculateMaxAbsDiff(root *Node) int {
	_, _, ad := dfsStep(root)

	return int(ad)
}

func dfsStep(node *Node) (float64, float64, float64) {
	if node.Left == nil && node.Right == nil {
		return float64(node.Val), float64(node.Val), 0
	}

	min := float64(node.Val)
	max := float64(node.Val)
	ad := 0.0

	if node.Left != nil {
		m, x, d := dfsStep(node.Left)
		tm := math.Min(m, float64(node.Val))
		tx := math.Max(x, float64(node.Val))
		ad = math.Max(math.Abs(tx-tm), math.Max(ad, d))

		min = math.Min(m, min)
		max = math.Max(x, max)
	}

	if node.Right != nil {
		m, x, d := dfsStep(node.Left)
		tm := math.Min(m, float64(node.Val))
		tx := math.Max(x, float64(node.Val))
		ad = math.Max(math.Abs(tx-tm), math.Max(ad, d))

		min = math.Min(m, min)
		max = math.Max(x, max)
	}

	return min, max, ad
}
