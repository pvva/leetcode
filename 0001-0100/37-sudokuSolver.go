package main

import (
	"fmt"
)

func main() {
	sudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(sudoku)
	solution := []string{
		"534678912",
		"672195348",
		"198342567",
		"859761423",
		"426853791",
		"713924856",
		"961537284",
		"287419635",
		"345286179",
	}
	for i, row := range sudoku {
		fmt.Println(string(row), solution[i])
	}
}

/*
https://leetcode.com/problems/sudoku-solver/

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
The '.' character indicates empty cells.

Example 1:

Input: board = [
["53",".",".","7",".",".",".","."],
["6",".",".","1","9","5",".",".","."],
[".","9","8",".",".",".",".","6","."],
["8",".",".",".","6",".",".",".","3"],
["4",".",".","8",".","3",".",".","1"],
["7",".",".",".","2",".",".",".","6"],
[".","6",".",".",".",".","2","8","."],
[".",".",".","4","1","9",".",".","5"],
[".",".",".",".","8",".",".","7","9"]]
Output: [
["5","3","4","6","7","8","9","1","2"],
["6","7","2","1","9","5","3","4","8"],
["1","9","8","3","4","2","5","6","7"],
["8","5","9","7","6","1","4","2","3"],
["4","2","6","8","5","3","7","9","1"],
["7","1","3","9","2","4","8","5","6"],
["9","6","1","5","3","7","2","8","4"],
["2","8","7","4","1","9","6","3","5"],
["3","4","5","2","8","6","1","7","9"]]

Explanation: The input board is shown above and the only valid solution.

Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit or '.'.
It is guaranteed that the input board has only one solution.
*/

func getPossibleNumbers(board [][]byte, x, y int) []int {
	// find appropriate value for row/column/3x3 cell
	res := []int{}
	invalidNumbers := [9]bool{}
	for i := 0; i < len(board[x]); i++ {
		if board[x][i] != '.' {
			invalidNumbers[int(board[x][i]-'1')] = true
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][y] != '.' {
			invalidNumbers[int(board[i][y]-'1')] = true
		}
	}
	_x := (x / 3) * 3
	_y := (y / 3) * 3
	for r := _x; r < _x+3; r++ {
		for c := _y; c < _y+3; c++ {
			if board[r][c] != '.' {
				invalidNumbers[int(board[r][c]-'1')] = true
			}
		}
	}

	for i, v := range invalidNumbers {
		if v {
			continue
		}
		res = append(res, i+1)
	}

	return res
}

func getNextXY(board [][]byte, x, y int) (int, int, bool) {
	if y < len(board[x])-1 {
		y++
	} else if x < len(board)-1 {
		y = 0
		x++
	} else {
		return 0, 0, true
	}
	return x, y, false
}

func solveSudokuBT(board [][]byte, x, y int) bool {
	if board[x][y] != '.' {
		nx, ny, stop := getNextXY(board, x, y)
		if stop {
			return true
		}

		return solveSudokuBT(board, nx, ny)
	}
	possibleNumbers := getPossibleNumbers(board, x, y)
	for _, v := range possibleNumbers {
		board[x][y] = byte('0' + v)
		nx, ny, stop := getNextXY(board, x, y)
		if stop || solveSudokuBT(board, nx, ny) {
			return true
		}
		board[x][y] = '.'
	}

	return false
}

func solveSudoku(board [][]byte) {
	solveSudokuBT(board, 0, 0)
}
