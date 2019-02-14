package main

import "fmt"

func main() {
	//fmt.Println(spiralOrder([][]int{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//}))
	fmt.Println(spiralOrder([][]int{
		{3},
		{2},
	}))
}

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
func spiralOrder(matrix [][]int) []int {
	y := len(matrix)
	if y == 0 {
		return []int{}
	}

	x := len(matrix[0])
	lOut := x * y
	out := make([]int, lOut)
	if x == 1 {
		for i := 0; i < y; i++ {
			out[i] = matrix[i][0]
		}

		return out
	}

	idx := 0
	yIdx := 0
	xIdx := 0
	yDir := 1
	xDir := 1
	isX := true
	xLim := x - 1
	yLim := y - 1

	for idx < lOut {
		out[idx] = matrix[yIdx][xIdx]
		if isX {
			xIdx += xDir
			if ((xDir == 1) && (xIdx >= xLim)) || ((xDir == -1) && (xIdx <= xLim)) {
				xIdx = xLim
				isX = false
				s := 1
				if xDir == 1 {
					xDir = -1
				} else {
					xDir = 1
					s = 2
				}
				xLim = x - xLim - s
			}
		} else {
			yIdx += yDir
			if ((yDir == 1) && (yIdx >= yLim)) || ((yDir == -1) && (yIdx <= yLim)) {
				yIdx = yLim
				isX = true
				s := 0
				if yDir == 1 {
					yDir = -1
				} else {
					yDir = 1
					s = 1
				}
				yLim = y - yLim - s
			}
		}
		idx++
	}

	return out
}
