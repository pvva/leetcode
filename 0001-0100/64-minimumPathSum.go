package main

func main() {
	println(7, minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	println(12, minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}

/*
https://leetcode.com/problems/minimum-path-sum/

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes
the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example 1:

Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

Example 2:

Input: grid = [[1,2,3],[4,5,6]]
Output: 12

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	buf := make([]int, len(grid[0]))
	for i := range buf {
		buf[i] = 1<<31 - 1
	}
	buf[0] = grid[0][0]

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if y < len(grid[x])-1 {
				sh := buf[y] + grid[x][y+1]
				buf[y+1] = min(buf[y+1], sh)
			}
			if x < len(grid)-1 {
				buf[y] += grid[x+1][y]
			}
		}
	}

	return buf[len(buf)-1]
}
