package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}), 16)
	fmt.Println(islandPerimeter([][]int{{1, 1}, {1, 1}}), 8)
}

/*
https://leetcode.com/problems/island-perimeter/

You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water,
and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell
is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Example:

Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Output: 16
*/

var dirs = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func islandPerimeter(grid [][]int) int {
	l := len(grid)
	if l == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid[0])

	sx, sy := -1, -1
outer:
	for x, row := range grid {
		for y, v := range row {
			if v == 1 {
				sx = x
				sy = y
				break outer
			}
		}
	}
	if sx == -1 {
		return 0
	}

	p := 0

	cLayer := map[[2]int]bool{{sx, sy}: true}
	for len(cLayer) > 0 {
		nLayer := make(map[[2]int]bool)

		for coords := range cLayer {
			grid[coords[0]][coords[1]] = -1
			for _, d := range dirs {
				cx := coords[0] + d[0]
				cy := coords[1] + d[1]

				if cx < 0 || cx == l || cy < 0 || cy == m {
					p++
				} else {
					if grid[cx][cy] == 0 {
						p++
					} else if grid[cx][cy] == 1 {
						nLayer[[2]int{cx, cy}] = true
					}
				}
			}
		}

		cLayer = nLayer
	}

	return p
}

func islandPerimeter2(grid [][]int) int {
	l := len(grid)
	if l == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid[0])

	p := 0
	for x, row := range grid {
		for y, v := range row {
			if v == 0 {
				continue
			}
			for _, d := range dirs {
				cx := x + d[0]
				cy := y + d[1]

				if cx < 0 || cx == l || cy < 0 || cy == m || grid[cx][cy] == 0 {
					p++
				}
			}
		}
	}

	return p
}

func islandPerimeter3(grid [][]int) int {
	l := len(grid)
	if l == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid[0])

	p := 2 * grid[0][0]
	for y := 1; y < m; y++ {
		p += grid[0][y] + (grid[0][y-1] ^ grid[0][y])
	}
	p += grid[0][m-1]

	for x := 1; x < l; x++ {
		p += grid[x][0] + (grid[x][0] ^ grid[x-1][0])
		for y := 1; y < m; y++ {
			p += (grid[x][y] ^ grid[x-1][y]) + (grid[x][y-1] ^ grid[x][y])
		}
		p += grid[x][m-1]
	}

	for y := 0; y < m; y++ {
		p += grid[l-1][y]
	}

	return p
}

func islandPerimeter4(grid [][]int) int {
	p := 0
	for x, row := range grid {
		for y, v := range row {
			if v == 0 {
				continue
			}
			p += 4
			if y > 0 && row[y-1] == 1 {
				p -= 2
			}
			if x > 0 && grid[x-1][y] == 1 {
				p -= 2
			}
		}
	}

	return p
}
