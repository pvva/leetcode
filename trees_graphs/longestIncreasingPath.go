package main

func main() {
	//println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	//println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	//println(longestIncreasingPath([][]int{{1}}))
	println(longestIncreasingPath([][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
		{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
		{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
		{79, 78, 77, 76, 75, 74, 73, 72, 71, 70},
		{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
		{99, 98, 97, 96, 95, 94, 93, 92, 91, 90},
		{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
		{119, 118, 117, 116, 115, 114, 113, 112, 111, 110},
		{120, 121, 122, 123, 124, 125, 126, 127, 128, 129},
		{139, 138, 137, 136, 135, 134, 133, 132, 131, 130},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}))
}

/*
Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down.
You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

Input: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].
*/
func longestIncreasingPath(matrix [][]int) int {
	// dfs from each cell + dynamic programming (remember results)
	if len(matrix) == 0 {
		return 0
	}
	max := 0
	lm := len(matrix)
	maxPaths := make([][]int, lm)
	for x := 0; x < len(matrix); x++ {
		maxPaths[x] = make([]int, len(matrix[x]))
	}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			cm := maxPathDfs(x, y, matrix, maxPaths)
			if max < cm {
				max = cm
			}
		}
	}

	return max
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func maxPathDfs(x, y int, matrix, maxPaths [][]int) int {
	if maxPaths[x][y] != 0 {
		return maxPaths[x][y]
	}

	max := 1
	for _, dir := range dirs {
		nx := x + dir[0]
		ny := y + dir[1]
		if nx < 0 || nx >= len(matrix) || ny < 0 || ny >= len(matrix[x]) || matrix[nx][ny] <= matrix[x][y] {
			continue
		}
		cm := 1 + maxPathDfs(nx, ny, matrix, maxPaths)
		if max < cm {
			max = cm
		}
	}
	maxPaths[x][y] = max

	return max
}
