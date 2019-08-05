package main

import "fmt"

/*
Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.
Input: [[1,1],[2,2],[3,3]]
Output: 3

Input: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
Output: 4
*/

type maxPointsOnALineTestCase struct {
	points [][]int
	result int
}

func main() {
	testCases := []maxPointsOnALineTestCase{
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
		{[][]int{{1, 1}, {1, 1}, {1, 1}}, 3},
		{[][]int{{1, 1}, {1, 1}, {2, 2}, {2, 2}}, 4},
		{[][]int{{1, 1}, {0, 1}, {2, 2}, {3, 3}, {3, 3}, {4, 4}, {4, 5}, {5, 5}}, 6},
	}

	for _, tc := range testCases {
		fmt.Println("Points: ", tc.points, ", expected: ", tc.result, ", result: ", maxPoints(tc.points))
	}
}

// use brute force iteration, n^2 - n steps
func maxPoints(points [][]int) int {
	l := len(points)
	if l < 3 {
		return l
	}

	maxP := 2

	for i := 0; i < l-1; i++ {
		samePoint := 1
		for j := i + 1; j < l; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				samePoint++
				continue
			}

			cMaxP := 2
			// points i and j are fixed, need to check other points if they form a line with i and j
			for p := 0; p < l; p++ {
				if p == i || p == j {
					continue
				}
				if isLine(points[i], points[j], points[p]) {
					cMaxP++
				}
			}

			if cMaxP > maxP {
				maxP = cMaxP
				println("0: new max is ", maxP)
			}
		}
		if samePoint > maxP {
			maxP = samePoint
			println("1: new max is ", maxP)
		}
	}

	return maxP
}

func isLine(p1, p2, p3 []int) bool {
	if p1[0] == p2[0] && p1[1] == p2[1] {
		return true
	}

	return int64(p3[0]-p1[0])*int64(p2[1]-p1[1]) == int64(p3[1]-p1[1])*int64(p2[0]-p1[0])
}
