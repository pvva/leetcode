package main

func main() {}

/*
Given a binary tree root. Split the binary tree into two subtrees by removing 1 edge such that the product of the sums
of the subtrees are maximized.

Since the answer may be too large, return it modulo 10^9 + 7.

Example 1:

Input: root = [1,2,3,4,5,6]
Output: 110

Example 2:

Input: root = [1,null,2,3,4,null,null,5,6]
Output: 90

Example 3:

Input: root = [2,3,9,10,7,8,6,5,4,11,1]
Output: 1025

Example 4:

Input: root = [1,1]
Output: 1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	sums := make(map[int]bool)

	total := calculateSums(root, sums)
	result := int64(0)

	for partialSum := range sums {
		otherSum := total - partialSum

		t := int64(otherSum) * int64(partialSum)
		if result < t {
			result = t
		}
	}

	return int(result % mod)
}

const mod = 1000000007

func calculateSums(root *TreeNode, sums map[int]bool) int {
	sum := root.Val
	if root.Left != nil {
		sum += calculateSums(root.Left, sums)
	}
	if root.Right != nil {
		sum += calculateSums(root.Right, sums)
	}

	sums[sum] = true

	return sum
}
