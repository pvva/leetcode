package main

func main() {
	println(maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}), 6)
}

/*
https://leetcode.com/problems/maximal-rectangle/

Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

Example:

Input:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
Output: 6
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	maxArea := 0
	rowHeights := make([]int, len(matrix[0]))
	for _, row := range matrix {
		for y, v := range row {
			if v == '1' {
				rowHeights[y]++
			} else {
				rowHeights[y] = 0
			}
		}

		sq := largestArea(rowHeights)
		if sq > maxArea {
			maxArea = sq
		}
	}

	return maxArea
}

func largestArea(heights []int) int {
	sq := 0

	mins := make([][2]int, len(heights)) // indices for left, right
	stack := make([]int, len(heights))
	si := -1

	for i, v := range heights {
		for si >= 0 && heights[stack[si]] >= v {
			si--
		}

		if si >= 0 {
			mins[i][0] = stack[si]
		} else {
			mins[i][0] = -1
		}
		si++
		stack[si] = i
	}

	si = -1
	for i := len(heights) - 1; i >= 0; i-- {
		for si >= 0 && heights[stack[si]] >= heights[i] {
			si--
		}

		if si >= 0 {
			mins[i][1] = stack[si]
		} else {
			mins[i][1] = len(heights)
		}
		si++
		stack[si] = i
	}

	for i, v := range heights {
		t := (mins[i][1] - mins[i][0] - 1) * v
		if t > sq {
			sq = t
		}
	}

	return sq
}
