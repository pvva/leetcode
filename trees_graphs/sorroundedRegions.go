package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		//{'X', 'X', 'X', 'X'},
		//{'X', 'O', 'O', 'X'},
		//{'X', 'X', 'O', 'X'},
		//{'X', 'O', 'X', 'X'},
		//{'X', 'X', 'X'},
		//{'X', 'O', 'X'},
		//{'X', 'X', 'X'},
		//{'X', 'X', 'X', 'X'},
		//{'O', 'X', 'O', 'X'},
		//{'X', 'O', 'X', 'O'},
		//{'O', 'X', 'O', 'X'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}

	solve(board)
	fmt.Println(board)
}

/*
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

X X X X
X O O X
X X O X
X O X X

After running your function, the board should be:
X X X X
X X X X
X X X X
X O X X
*/

func solve(board [][]byte) {
	lx := len(board) - 1
	for x, row := range board {
		for y := range row {
			if (x == 0 || x == lx || y == 0 || y == len(row)-1) && board[x][y] == 'O' {
				markPointAsUnbounded(x, y, board)
			}
		}
	}
	for x, row := range board {
		for y := range row {
			if board[x][y] == 'U' {
				// this is unbounded, set to O again
				board[x][y] = 'O'
			} else if board[x][y] == 'O' {
				// this is inside bounds of X, set to X
				board[x][y] = 'X'
			}
		}
	}
}

var fillDirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func markPointAsUnbounded(x, y int, board [][]byte) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[x]) {
		return
	}

	if board[x][y] == 'O' {
		board[x][y] = 'U'

		for _, dir := range fillDirs {
			markPointAsUnbounded(x+dir[1], y+dir[0], board)
		}
	}
}
