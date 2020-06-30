package main

func main() {
	println(uniquePaths(3, 2), 3)
	println(uniquePaths(7, 3), 28)
}

/*
https://leetcode.com/problems/unique-paths/

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right
corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?

Example 1:

Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
Example 2:

Input: m = 7, n = 3
Output: 28
*/

func uniquePaths(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}

	pathsCount := make([]int, m)
	for y := 0; y < m; y++ {
		pathsCount[y] = 1
	}

	for x := 1; x < n; x++ {
		for y := 1; y < m; y++ {
			pathsCount[y] += pathsCount[y-1]
		}
	}

	return pathsCount[m-1]
}
