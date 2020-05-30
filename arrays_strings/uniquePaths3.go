package main

func main() {
	println(uniquePathsIII([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}), 2)
}

/*
https://leetcode.com/problems/unique-paths-iii/

On a 2-dimensional grid, there are 4 types of squares:

1 represents the starting square.  There is exactly one starting square.
2 represents the ending square.  There is exactly one ending square.
0 represents empty squares we can walk over.
-1 represents obstacles that we cannot walk over.

Return the number of 4-directional walks from the starting square to the ending square,
that walk over every non-obstacle square exactly once.

Example 1:

Input: [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
Output: 2
Explanation: We have the following two paths:
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
*/

func uniquePathsIII(grid [][]int) int {
	maxSteps, sx, sy := calculateMaxStepsAndStart(grid)
	if maxSteps == 0 {
		return 0
	}

	return dfsStep(grid, sx, sy, 1, maxSteps)
}

func calculateMaxStepsAndStart(grid [][]int) (int, int, int) {
	sx := -1
	sy := -1
	steps := 0

	for x, row := range grid {
		for y, v := range row {
			if v == -1 {
				continue
			}
			steps++
			if v == 1 {
				sx = x
				sy = y
			}
		}
	}

	return steps, sx, sy
}

func dfsStep(grid [][]int, x, y, currentSteps, maxSteps int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == -1 ||
		(grid[x][y] == 2 && currentSteps < maxSteps) {
		return 0
	}

	if grid[x][y] == 2 && currentSteps == maxSteps {
		return 1
	}

	v := grid[x][y]
	grid[x][y] = -1

	numCompletePathsFound := dfsStep(grid, x, y+1, currentSteps+1, maxSteps)
	numCompletePathsFound += dfsStep(grid, x+1, y, currentSteps+1, maxSteps)
	numCompletePathsFound += dfsStep(grid, x, y-1, currentSteps+1, maxSteps)
	numCompletePathsFound += dfsStep(grid, x-1, y, currentSteps+1, maxSteps)

	grid[x][y] = v

	return numCompletePathsFound
}
