package main

func main() {
	println(findPaths(2, 2, 2, 0, 0), 6)
	println(findPaths(1, 3, 3, 0, 1), 12)
}

/*
https://leetcode.com/problems/out-of-boundary-paths/

There is an m by n grid with a ball. Given the start coordinate (i,j) of the ball, you can move the ball to
adjacent cell or cross the grid boundary in four directions (up, down, left, right). However, you can at most
move N times. Find out the number of paths to move the ball out of grid boundary. The answer may be very
large, return it after mod 10^9 + 7.

Example 1:
Input: m = 2, n = 2, N = 2, i = 0, j = 0
Output: 6

Example 2:
Input: m = 1, n = 3, N = 3, i = 0, j = 1
Output: 12
*/

func findPaths(m int, n int, N int, i int, j int) int {
	if N == 0 {
		return 0
	}

	pathsCount := make([][][]int, N)
	for k := range pathsCount {
		pathsCount[k] = make([][]int, m)
		for x := range pathsCount[k] {
			pathsCount[k][x] = make([]int, n)
			for y := range pathsCount[k][x] {
				pathsCount[k][x][y] = -1
			}
		}
	}

	return findPathsWithMemory(m, n, N, i, j, pathsCount)
}

const mod = 1000000007

func findPathsWithMemory(m, n, N, ci, cj int, pathsCount [][][]int) int {
	if ci < 0 || cj < 0 || ci == m || cj == n {
		return 1
	}
	if N == 0 {
		return 0
	}

	t := pathsCount[N-1][ci][cj]
	if t != -1 {
		return t
	}

	result := findPathsWithMemory(m, n, N-1, ci-1, cj, pathsCount) % mod
	result = (result + findPathsWithMemory(m, n, N-1, ci, cj+1, pathsCount)) % mod
	result = (result + findPathsWithMemory(m, n, N-1, ci+1, cj, pathsCount)) % mod
	result = (result + findPathsWithMemory(m, n, N-1, ci, cj-1, pathsCount)) % mod

	pathsCount[N-1][ci][cj] = result

	return result
}
