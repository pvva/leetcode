package main

import (
	"fmt"
	"strings"
)

func main() {
	printResults(solveNQueens(5))
}

func printResults(r [][]string) {
	for _, row := range r {
		fmt.Println(strings.Join(row, "\n"))
		fmt.Println()
	}
}

/*
https://leetcode.com/problems/n-queens/

The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a
queen and an empty space, respectively.

Example 1:

Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

Example 2:

Input: n = 1
Output: [["Q"]]

Constraints:

1 <= n <= 9
*/

// this can be optimized
func updateCounts(counts [][]int, dir, n, ir, ic int) {
	for i := 0; i < n; i++ {
		counts[ir][i] += dir
		if ir != i || ic != i {
			counts[i][ic] += dir
		}
		if i > 0 {
			for _, shifts := range [][]int{{1, 1}, {-1, -1}, {-1, 1}, {1, -1}} {
				nir := ir + shifts[0]*i
				nic := ic + shifts[1]*i
				if nir < n && nic < n && nir >= 0 && nic >= 0 {
					counts[nir][nic] += dir
				}
			}
		}
	}

}

func solveNQueensBT(n int, board [][]byte, fr int, counts [][]int, res *[][]string) {
	if n == 0 {
		t := make([]string, len(board))
		for i, r := range board {
			t[i] = string(r)
		}
		*res = append(*res, t)
		return
	}
	for ir := fr; ir < len(board); ir++ {
		for ic := range board[ir] {
			if counts[ir][ic] > 0 {
				continue
			}
			board[ir][ic] = 'Q'
			updateCounts(counts, 1, len(board), ir, ic)
			solveNQueensBT(n-1, board, ir+1, counts, res)
			board[ir][ic] = '.'
			updateCounts(counts, -1, len(board), ir, ic)
		}
	}
}

func solveNQueens(n int) [][]string {
	res := [][]string{}

	for j := 0; j < n; j++ {
		board := make([][]byte, n)
		counts := make([][]int, n)
		for i := range board {
			board[i] = make([]byte, n)
			counts[i] = make([]int, n)
			for k := range board[i] {
				board[i][k] = '.'
			}
		}

		board[0][j] = 'Q'
		updateCounts(counts, 1, n, 0, j)
		solveNQueensBT(n-1, board, 1, counts, &res)
	}

	return res
}
