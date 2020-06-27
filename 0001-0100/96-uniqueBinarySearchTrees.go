package main

func main() {
	println(numTrees(3), 5)
	println(numTrees(4))
}

/*
https://leetcode.com/problems/unique-binary-search-trees/

Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

Example:

Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

func numTrees(n int) int {
	return numTreesRec(n, map[int]int{0: 1, 1: 1, 2: 2})
}

func numTreesRec(n int, memo map[int]int) int {
	if v, ex := memo[n]; ex {
		return v
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		cnt += numTreesRec(i-1, memo) * numTreesRec(n-i, memo)
	}
	memo[n] = cnt

	return cnt
}
