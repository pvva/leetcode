package main

func main() {
	println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}), 2)
}

/*
https://leetcode.com/problems/unique-paths-ii/

A robot is located at the top-left corner of a m x n grid.
The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid.

Now consider if some obstacles are added to the grids. How many unique paths would there be?
An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	l := len(obstacleGrid[0])
	if l == 0 {
		return 0
	}

	pathsCount := make([][]int, len(obstacleGrid))
	for i := range pathsCount {
		pathsCount[i] = make([]int, l)
		if i == 0 {
			if obstacleGrid[0][0] != 1 {
				pathsCount[0][0] = 1
			}

			for j := 1; j < l; j++ {
				if obstacleGrid[i][j] != 1 {
					pathsCount[i][j] = pathsCount[i][j-1]
				}
			}
		} else {
			if obstacleGrid[i][0] != 1 {
				pathsCount[i][0] = pathsCount[i-1][0]
			}
		}
	}

	for i := 1; i < len(pathsCount); i++ {
		for j := 1; j < l; j++ {
			if obstacleGrid[i][j] != 1 {
				pathsCount[i][j] = pathsCount[i-1][j] + pathsCount[i][j-1]
			}
		}
	}

	return pathsCount[len(obstacleGrid)-1][l-1]
}
