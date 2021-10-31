package main

func main() {
	println(true, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
}

/*
https://leetcode.com/problems/search-a-2d-matrix/

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.

Example 1:

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/

func searchMatrix(matrix [][]int, target int) bool {
	h := len(matrix) - 1
	l := 0
	li := len(matrix[0]) - 1
	var row []int
	for l <= h {
		m := (h + l) / 2
		row = matrix[m]
		if target >= matrix[m][0] && target <= matrix[m][li] {
			break
		}
		if target < matrix[m][0] {
			h = m - 1
		} else if target > matrix[m][li] {
			l = m + 1
		}
	}
	l = 0
	h = li
	for l <= h {
		m := (h + l) / 2
		if row[m] == target {
			return true
		}
		if target < row[m] {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}
